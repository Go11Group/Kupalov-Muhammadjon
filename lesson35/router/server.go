package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

func CreateServer(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/leetcode.uz")
	r.HandleFunc("/users/", handler.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", handler.GetUserByID).Methods("GET")
	r.HandleFunc("/create-user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/update-user/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete-user/{id}", handler.DeleteUser).Methods("DELETE")

	// r.HandleFunc("/problems", handler.GetProblems).Methods("GET")
	// r.HandleFunc("/problems/{id}", handler.GetProblemByID).Methods("GET")
	// r.HandleFunc("/problems", handler.CreateProblem).Methods("POST")
	// r.HandleFunc("/problems/{id}", handler.UpdateProblem).Methods("PUT")
	// r.HandleFunc("/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	// r.HandleFunc("/submissions", handler.GetSubmissions).Methods("GET")
	// r.HandleFunc("/submissions/{id}", handler.GetSubmissionByID).Methods("GET")
	// r.HandleFunc("/submissions", handler.CreateSubmission).Methods("POST")
	// r.HandleFunc("/submissions/{id}", handler.UpdateSubmission).Methods("PUT")
	// r.HandleFunc("/submissions/{id}", handler.DeleteSubmission).Methods("DELETE")

	return r
}
