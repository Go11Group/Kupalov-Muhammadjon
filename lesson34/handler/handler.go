package handler

import (
	"module/storage/postgres"
	"net/http"
)

type Handler struct{
	UserRepo *postgres.UserRepo
	ProductRepo *postgres.ProductRepo
	UserProductRepo *postgres.UserProductRepo
}

func NewHandler(u *postgres.UserRepo, p *postgres.ProductRepo, up *postgres.UserProductRepo) *Handler{
	return &Handler{u, p, up}
}

func CreateServer(handler *Handler) *http.Server{
	mux := http.NewServeMux()
	//User CRUD
	mux.HandleFunc("POST /create-user", handler.CreateUser)
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("PUT /update-user", handler.UpdateUser)
	mux.HandleFunc("DELETE /delete-user", handler.DeleteUser)

	// Product CRUD
	mux.HandleFunc("POST /create-product", handler.CreateProduct)
	mux.HandleFunc("GET /products", handler.GetProducts)
	mux.HandleFunc("PUT /update-product", handler.UpdateProduct)
	mux.HandleFunc("DELETE /delete-product", handler.DeleteProduct)

	// User Product CRUD
	mux.HandleFunc("POST /create-user-product", handler.CreateUserProduct)
	mux.HandleFunc("GET /user-products", handler.GetUserProducts)
	mux.HandleFunc("PUT /update-user-product", handler.UpdateUserProduct)
	mux.HandleFunc("DELETE /delete-user-product", handler.DeleteUserProduct)



	return &http.Server{Addr: ":8080", Handler: mux}
}