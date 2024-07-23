package main

import (
	"lesson64/api"
	"lesson64/storage/redis"
	"log"
)

func main() {

	db, err := redis.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := api.NewRouter(db)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
