package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

func CreateServer(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/leetcode.uz").Subrouter()

	// User CRUD
	api.HandleFunc("/users/", handler.GetUsers).Methods("GET")
	api.HandleFunc("/user/{id}", handler.GetUserByID).Methods("GET")
	api.HandleFunc("/create-user", handler.CreateUser).Methods("POST")
	api.HandleFunc("/update-user/{id}", handler.UpdateUser).Methods("PUT")
	api.HandleFunc("/delete-user/{id}", handler.DeleteUser).Methods("DELETE")

	// Problem CRUD
	api.HandleFunc("/problems", handler.GetProblems).Methods("GET")
	api.HandleFunc("/problem/{id}", handler.GetProblemByID).Methods("GET")
	api.HandleFunc("/problem", handler.CreateProblem).Methods("POST")
	api.HandleFunc("/problems/{id}", handler.UpdateProblem).Methods("PUT")
	api.HandleFunc("/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	// Language CRUD
	api.HandleFunc("/languages", handler.GetLanguages).Methods("GET")
	api.HandleFunc("/language/{id}", handler.GetLanguageByID).Methods("GET")
	api.HandleFunc("/create-language", handler.CreateLanguage).Methods("POST")
	api.HandleFunc("/update-language/{id}", handler.UpdateLanguage).Methods("PUT")
	api.HandleFunc("/delete-language/{id}", handler.DeleteLanguage).Methods("DELETE")

	// Topic CRUD
	api.HandleFunc("/topics", handler.GetTopics).Methods("GET")
	api.HandleFunc("/topic/{id}", handler.GetTopicByID).Methods("GET")
	api.HandleFunc("/create-topic", handler.CreateTopic).Methods("POST")
	api.HandleFunc("/update-topic/{id}", handler.UpdateTopic).Methods("PUT")
	api.HandleFunc("/delete-topic/{id}", handler.DeleteTopic).Methods("DELETE")

	// TopicProblem CRUD
	api.HandleFunc("/problemtopics", handler.GetTopics).Methods("GET")
	api.HandleFunc("/problemtopic/{id}", handler.GetTopicByID).Methods("GET")
	api.HandleFunc("/create-problemtopic", handler.CreateTopic).Methods("POST")
	api.HandleFunc("/update-problemtopic/{id}", handler.UpdateTopic).Methods("PUT")
	api.HandleFunc("/delete-problemtopic/{id}", handler.DeleteTopic).Methods("DELETE")

	// Submission CRUD
	// api.HandleFunc("/submissions", handler.GetSubmissions).Methods("GET")
	// api.HandleFunc("/submissions/{id}", handler.GetSubmissionByID).Methods("GET")
	// api.HandleFunc("/submissions", handler.CreateSubmission).Methods("POST")
	// api.HandleFunc("/submissions/{id}", handler.UpdateSubmission).Methods("PUT")
	// api.HandleFunc("/submissions/{id}", handler.DeleteSubmission).Methods("DELETE")

	return r
}
