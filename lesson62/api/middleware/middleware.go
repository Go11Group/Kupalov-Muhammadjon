package middleware

import (
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct{
	enforcer *casbin.Enforcer
}

func (c *casbinPermission) GetRole(ctx *gin.Context) (string, int){
	role := ctx.GetHeader("Authorization")

	if role == "" {
		return "unauthorized", http.StatusUnauthorized
	}
	return role, 0
}

func (c *casbinPermission) CheckPermission(ctx *gin.Context) (bool, error){
	subject, status := c.GetRole(ctx)
	if status != 0{
		return false, errors.New("error while getting a role")
	}
	action := ctx.Request.Method
	object := ctx.Request.URL

	allow, err := c.enforcer.Enforce(subject, object.String(), action)
	if err != nil {
		return false, err
	}

	return allow, nil
}


func PermissonMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(ctx *gin.Context) {
		res, err := casbHandler.CheckPermission(ctx)

		if err != nil {
			ctx.AbortWithError(500, err)
		}
		if !res{
			ctx.AbortWithStatusJSON(401, "unauthorized")
		}
		ctx.Next()
	}
}