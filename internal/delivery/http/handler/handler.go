package handler

import (
	"fibonachi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(jsonContentTypeMiddleware)

	api := router.Group("/api")
	{
		fibonacci := api.Group("/fibonacci")
		{
			fibonacci.POST("/", h.Calculate)
		}
	}
	return router
}
