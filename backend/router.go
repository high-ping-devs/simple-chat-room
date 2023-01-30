package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"github.com/high-ping-devs/simple-chat-room/backend/controllers"
)

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/health", health)

	public := r.Group("/auth")
	{
		public.POST("/login", controllers.Login)
		public.POST("/register", controllers.Register)
	}

	manageAuth := r.Group("/auth")
	manageAuth.Use(auth.Middleware())
	{
		manageAuth.POST("/logout", controllers.Logout)
		manageAuth.POST("/refresh", controllers.Refresh)
	}

	ws := r.Group("/ws")
	ws.Use(auth.Middleware())
	{
		ws.Handle("GET", "/chat", chatHandler)
	}

	return r
}

func health(c *gin.Context) {
	c.Status(200)
}

func chatHandler(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)

	c.JSON(200, gin.H{
		"message": "pong",
		"claims":  claims,
	})
}
