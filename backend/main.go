package main

import (
	"github.com/high-ping-devs/simple-chat-room/backend/database"
	"github.com/high-ping-devs/simple-chat-room/backend/session"
)

func main() {
	var s session.Storage
	database.Connect()
	database.Migrate()
	s.Create()
}
