package main

import (
	"lesson62/api"
	"lesson62/models"
	"lesson62/pkg/logger"
	"lesson62/storage/redis"

	"go.uber.org/zap"
)

func main() {
	logger, err := logger.New("app.log")
	if err != nil {
		panic(err)
	}
	db, err := redis.ConnectDB()
	if err != nil {
		panic(err)
	}
	systemConfig := &models.SystemConfig{
		Logger: logger,
		RedisDb: db,
	}

	router := api.NewRouter(systemConfig)

	err = router.Run(":8080")
	if err != nil {
		logger.Fatal("gin is not working not runnning ", zap.Error(err))
	}
}
