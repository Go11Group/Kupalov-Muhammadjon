package main

import (
	"log"
	"module/storage/postgres"
	"net/http"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	http.HandleFunc("GET /getStudents", st.GetStudents)
	http.HandleFunc("POST /createStudent", st.CreateStudent)
	http.HandleFunc("DELETE /deleteStudent", st.DeleteStudent)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
