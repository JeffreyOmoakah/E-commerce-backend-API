package products

import (
	"net/http"

	"github.com/JeffreyOmoakah/E-commerce-backend-API/internal/json"
)

type handler struct{
	service Service
}

func NewHandler (service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter,r  *http.Request){
	// 1. Call the service -> list the products
   	err := h.service.ListProduct(r.Context())
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// 2. Return the JSON in an HTTP response
	products := struct {
	products []string `json:"products"`
	}{}
	
	json.Write(w, http.StatusOK, products)
}