package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dyhalmeida/go-apis/internal/dto"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB        database.UserDbInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(db database.UserDbInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Create user godoc
// @Summary Create User
// @Description Create User: A user can be created to access the system by passing their email, name and password.
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserInputDTO true "user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {

	var userDTO dto.UserInputDTO

	err := json.NewDecoder(req.Body).Decode(&userDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	res.WriteHeader(http.StatusCreated)

}

// GetJWTToken godoc
// @Summary Get a access token
// @Description Get a access token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CredentialsInputDTO true "user credentials"
// @Success 200 {object} dto.CredentialsOutputDTO
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users/token [post]
func (h *UserHandler) GetJwtToken(res http.ResponseWriter, req *http.Request) {

	jwt := req.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := req.Context().Value("jwtExpiresIn").(int)

	var credentialsDTO dto.CredentialsInputDTO
	err := json.NewDecoder(req.Body).Decode(&credentialsDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(credentialsDTO.Email)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	if !user.IsValidPassword(credentialsDTO.Password) {
		res.WriteHeader(http.StatusUnauthorized)
		errorMessage := Error{
			Message: errors.New("Unauthorized").Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.CredentialsOutputDTO{
		AccessToken: tokenString,
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(accessToken)
}
