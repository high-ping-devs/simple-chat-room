package main

import (
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/health", health)

	// auth := r.Group("/auth")
	// {
	// 	auth.post("/login", controllers.login)
	// 	auth.post("/register", controllers.register)
	// 	auth.post("/logout", controllers.logout)
	// 	auth.post("/refresh", controllers.refresh)
	// }

	// ws := r.Group("/ws")
	// ws.Use(auth.Middleware)
	// {
	// 	ws.Handle("/chat", chatHandler)
	// }

	return r
}

func health(c *gin.Context) {
	c.Status(200)
}
