package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"
)

// Create
func (h *Handler) CreateUserProduct(w http.ResponseWriter, r *http.Request) {
	newUserProduct := model.UserProducts{}
	err := json.NewDecoder(r.Body).Decode(&newUserProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decode %s", err)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = h.UserProductRepo.CreateUserProduct(newUserProduct.ProductId, newUserProduct.UserId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while writing user %s", err)))
		return
	}
	w.Write([]byte("Success"))

}

// Read
func (h *Handler) GetUserProducts(w http.ResponseWriter, r *http.Request){
	filter := model.FilterUserProducts{}
	w.Write([]byte(r.URL.Query().Encode()))
	q := r.URL.Query()
	if q.Has("id"){
		id, err := strconv.Atoi(q.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while getting filter user %s", err)))
			return
		}
		filter.Id = &id
	}
	if q.Has("user_id"){
		user_id, err := strconv.Atoi(q.Get("user_id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while getting filter user %s", err)))
			return
		}
		filter.Id = &user_id
	}
	if q.Has("product_id"){
		product_id, err := strconv.Atoi(q.Get("product_id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error while getting filter user %s", err)))
			return
		}
		filter.Id = &product_id
	}
	
	userPr, err := h.UserProductRepo.GetUserProduct(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while getting users %s", err)))
		return
	}
	err = json.NewEncoder(w).Encode(userPr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while encoding users %s", err)))
		return
	}	

	w.Write([]byte("Success"))
}
// Update
func (h *Handler) UpdateUserProduct(w http.ResponseWriter, r *http.Request){
	up := model.UserProducts{}
	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decoding %s", err)))
		return
	}
	err = h.UserProductRepo.UpdateUserProduct(up)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while updating user %s", err)))
		return
	}

	w.Write([]byte("Success"))
}
// Delete
func (h *Handler) DeleteUserProduct(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	if !q.Has("id"){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("send id of the user "))
		return
	}
	id, err := strconv.Atoi(q.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while coverting id %s", err)))
		return
	}
	err = h.UserProductRepo.DeleteUserProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while deleting user %s", err)))
		return
	}
	
	w.Write([]byte("Success"))
}