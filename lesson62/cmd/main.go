package main

import (
	"lesson62/api"
	"lesson62/models"
	"lesson62/pkg/logger"
	"lesson62/storage/redis"

	"github.com/casbin/casbin/v2"
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
	casbinEnforcer, err := casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")
	if err != nil {
		logger.Error("Error while loading model and policy")
		return
	}
	systemConfig := &models.SystemConfig{
		Logger: logger,
		RedisDb: db,
		CasbinEnforcer: casbinEnforcer,
	}

	router := api.NewRouter(systemConfig)

	err = router.Run(":8080")
	if err != nil {
		logger.Fatal("gin is not working not runnning ", zap.Error(err))
	}
}
