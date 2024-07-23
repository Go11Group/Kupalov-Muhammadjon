package main

import (
	"context"
	"fmt"
	"lesson64/storage/redis"
	"log"
	// "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb, err := redis.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	st := redis.NewStockRepo(rdb)

	pubsub := rdb.Subscribe(ctx, "apple", "tesla", "amazon")
	defer pubsub.Close()

	_, err = pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		err = st.WritePrice(ctx, msg.Payload)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
	}
}
