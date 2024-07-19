package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	connStr := "user=postgres password=root dbname=listenup_user_service sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Assume file is uploaded to a storage service and we have the file URL
	fileURL := "https://your-storage-service.com/path/to/audio/file.mp3"

	// Insert metadata into database
	query := `INSERT INTO episodes (podcast_id, user_id, title, description, file_url) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, 1, 1, "Episode 1", "Description of episode 1", fileURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Metadata inserted successfully")
}
