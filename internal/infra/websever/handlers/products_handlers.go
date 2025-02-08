package handlers

import (
	"encoding/json"
	"net/http"

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

func (h *ProductHandler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var productInputDTO dto.ProductInputDTO
	err := json.NewDecoder(req.Body).Decode(&productInputDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(productInputDTO.Name, productInputDTO.Price)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(product)

}

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
