package handler

import (
	"url-shortener/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func New_Handler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("send_url", h.send_url)
		api.GET("get_url/:short_url", h.get_url)
	}

	return router
}
