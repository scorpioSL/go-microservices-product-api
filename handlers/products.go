package handlers

import (
	"PRODUCT-API/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.logger.Printf("Producst API")
	if r.Method == http.MethodGet {
		handleGetProducts(rw, r)
		return
	}
}

func handleGetProducts(rw http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(data.GetProducts())
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Something went wrong!"))
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write((json))
}
