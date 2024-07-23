package main

import (
	"context"
	"encoding/json"
	"lesson64/models"
	"math/rand"
	"sync"
	"time"

	"log"

	"github.com/redis/go-redis/v9"
)

func main() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{

		Addr: "localhost:6379",
	})

	wait := sync.WaitGroup{}

	for {
		wait.Wait()
		wait.Add(1)
		go publishPrices(ctx, rdb, &wait)
	}

}

func publishPrices(ctx context.Context, rdb *redis.Client, wg *sync.WaitGroup) {

	price := rand.Intn(1000)
	info := models.Stock{
		CompanyName: "Apple",
		Price:       price,
		CurrentTime: time.Now(),
	}
	data, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	err = rdb.Publish(ctx, "apple", data).Err()
	if err != nil {
		log.Fatal(err)
	}

	price = rand.Intn(1000)
	info = models.Stock{
		CompanyName: "Tesla",
		Price:       price,
		CurrentTime: time.Now(),
	}
	data, err = json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(ctx, "tesla", data).Err()
	if err != nil {

		log.Fatal(err)

	}

	price = rand.Intn(1000)
	info = models.Stock{
		CompanyName: "Amazon",
		Price:       price,
		CurrentTime: time.Now(),
	}
	data, err = json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(ctx, "amazon", data).Err()
	if err != nil {

		log.Fatal(err)

	}

	time.Sleep(time.Millisecond * 200)
	wg.Done()
}
