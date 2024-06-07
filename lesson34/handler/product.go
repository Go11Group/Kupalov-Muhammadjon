package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"
)

// Create
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request){
	newProduct := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decoding %s ", err)))
		return
	}
	err = h.ProductRepo.CreateProduct(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while inserting user %s ", err)))
		return
	}
	w.Write([]byte("Success"))
}
// Read
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	filter := model.FilterProducts{}
	if q.Has("id"){
		id, err := strconv.Atoi(q.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while converting id %s ", err)))
			return
		}
		filter.Id = &id
	}
	if q.Has("name"){
		name := q.Get("name")
		filter.Name = &name
	}
	if q.Has("description"){
		description := q.Get("description")
		filter.Description = &description
	}
	if q.Has("price"){
		price, err := strconv.ParseFloat(q.Get("price"), 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while converting price %s ", err)))
			return
		}
		p32 := float32(price)
		filter.Price = &p32
	}
	if q.Has("stock_quantity"){
		stockQuantity, err := strconv.Atoi(q.Get("stock_quantity"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while converting stock_quantity %s ", err)))
			return
		}
		filter.StockQuantity = &stockQuantity
	}
	fmt.Println(filter)
	products, err := h.ProductRepo.GetProduct(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while getting products %s ", err)))
		return
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while encoding products %s ", err)))
		return
	}

	w.Write([]byte("Success"))
}
// Update
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	product := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decoding product %s ", err)))
		return
	}
	err = h.ProductRepo.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while updating product %s ", err)))
		return
	}
	w.Write([]byte("Success"))
}
// Delete
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	if !q.Has("id"){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error no id"))
		return
	}
	id, err := strconv.Atoi(q.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while converting id %s ", err)))
		return
	}
	err = h.ProductRepo.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while deleting product %s ", err)))
		return
	}
	w.Write([]byte("Success"))
}
