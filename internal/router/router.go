package router

import (
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

	api := r.Group("/api/v1")

	// публичные роуты
	auth := api.Group("/auth")
	{
		auth.POST("/login", router.login)
		auth.POST("/register", router.register)
	}

	// защищённые роуты
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", router.profile)
		// protected.POST("/books", ...)
	}

	return r
}

func (router *Router) Run() error {
	r := router.SetupRouter()
	return r.Run()
}
