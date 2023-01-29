package main

import (
	"context"
	"log"
	"time"

	"github.com/high-ping-devs/simple-chat-room/backend/sessionStorage"
)

func main() {
	var s sessionStorage.S
	s.Create()

	if err := s.Set(context.Background(), "test", "test"); err != nil {
		log.Panic(err)
	}

	for {
		if value, err := s.Get(context.Background(), "test"); err != nil {
			log.Panic(err)
		} else {
			log.Println(value)
		}

		time.Sleep(5 * time.Second)
	}
}
