package models

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"

	"go.uber.org/zap"
)

type SystemConfig struct {
	Logger         *zap.Logger
	RedisDb        *redis.Client
	CasbinEnforcer *casbin.Enforcer
}
