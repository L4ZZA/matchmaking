package handlers

import (
	"log"
	"net/http"
	"strconv"

	"example.com/data"
	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		p.l.Println("GetProducts - Failed to parse to JSON", lp, err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println("GetProducts - COMPLETED")
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("AddProduct - ERROR parsing JSON", prod, err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	p.l.Println("AddProduct - COMPLETED")
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		p.l.Println("UpdateProducts - can't convert id", id, err)
		http.Error(rw, "Unable to cast id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("UpdateProducts - Error Parsing from JSON | ", prod, err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		p.l.Println("UpdateProducts - ERROR2 ", err)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("UpdateProducts - ERROR3")
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
	p.l.Println("UpdateProducts - COMPLETED")
}
