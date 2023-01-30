package main

import (
	"log"
	"os"
	"strings"

	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"github.com/high-ping-devs/simple-chat-room/backend/database"
)

var (
	port = strings.TrimSpace(os.Getenv("BACKEND_PORT"))
)

func main() {
	log.Println("⏳ Starting server...")

	var wl auth.WhiteList
	database.Connect()
	database.Migrate()

	wl.Create()
	log.Println("🗃️ Redis connected successfully")

	log.Println("🚀 Server started on port " + port)
	router().Run(":" + port)
}
