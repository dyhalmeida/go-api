package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dyhalmeida/go-apis/internal/dto"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
	entityPkg "github.com/dyhalmeida/go-apis/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductDbInterface
}

func NewProductHandler(db database.ProductDbInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create Product
// @Description Create Products
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.ProductInputDTO true "product request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var productInputDTO dto.ProductInputDTO
	err := json.NewDecoder(req.Body).Decode(&productInputDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}
	product, err := entity.NewProduct(productInputDTO.Name, productInputDTO.Price)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}
	err = h.ProductDB.Create(product)
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

// Get Product godoc
// @Summary Get Product
// @Description Get Product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "id" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: errors.New("bad request").Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(product)

}

// Update Product godoc
// @Summary Update Product
// @Description Update Product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "id" Format(uuid)
// @Param request body dto.ProductInputDTO true "product"
// @Success 200
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(req.Body).Decode(&product)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(product.ID.String())

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPkg.ParseID(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)

}

// List Products godoc
// @Summary List Products
// @Description List Products
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Param sort query string false "sort"
// @Success 200 {array} entity.Product
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(res http.ResponseWriter, req *http.Request) {
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	sort := req.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(products)

}
