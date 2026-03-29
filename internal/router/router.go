package router

import (
	"github.com/trmyxx/AuthService/internal/middleware"
	"github.com/trmyxx/AuthService/internal/service"
	"github.com/trmyxx/AuthService/internal/storage"

	"github.com/gin-gonic/gin"
)

type Router struct {
	storage storage.Storage
	service service.Service
}

func NewRouter(storage storage.Storage, service service.Service) *Router {
	return &Router{
		storage: storage,
		service: service,
	}
}

func (router *Router) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/signup", router.Signup)
	r.POST("api/v1/login", router.Login)
	r.GET("/api/v1/validate", middleware.RequireAuth, router.Validate)

	return r
}

func (router *Router) Run() error {
	r := router.SetupRouter()
	return r.Run()
}
