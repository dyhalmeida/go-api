package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dyhalmeida/go-apis/internal/dto"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
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
