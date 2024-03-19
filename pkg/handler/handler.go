package handler

import (
    "login_api/pkg/service"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type Handler struct {
    services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.New()

    // Configure CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Allow requests from this origin
        AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // Allow these HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type"}, // Allow these request headers
        AllowCredentials: true, // Allow sending cookies
    }))

    auth := router.Group("/auth")
    {
        auth.POST("/sign-up", h.signUp)
        auth.POST("/sign-in", h.signIn)
    }
    return router
}
