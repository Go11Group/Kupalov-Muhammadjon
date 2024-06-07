package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"
)

// Create
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := model.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decode %s", err)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = h.UserRepo.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while writng user %s", err)))
		return
	}
	w.Write([]byte("Success"))

}

// Read
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request){
	filter := model.FilterUser{}
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
	if q.Has("username"){
		username := q.Get("username")
		filter.Username = &username
	}
	users, err := h.UserRepo.GetUser(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while getting users %s", err)))
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while encoding users %s", err)))
		return
	}	

	w.Write([]byte("Success"))
}
// Update
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request){
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error while decoding %s", err)))
		return
	}
	err = h.UserRepo.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while updating user %s", err)))
		return
	}

	w.Write([]byte("Success"))
}
// Delete
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request){
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
	err = h.UserRepo.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while deleting user %s", err)))
		return
	}

	w.Write([]byte("Success"))
}