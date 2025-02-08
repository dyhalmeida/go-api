package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dyhalmeida/go-apis/internal/dto"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserDbInterface
}

func NewUserHandler(db database.UserDbInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

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
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)

}
