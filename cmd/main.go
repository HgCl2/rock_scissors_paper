package main

import (
	"log"

	"github.com/rock_scissors_paper/handler"
)

func main() {
	srv := new(handler.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error run http server %s", err.Error())
	}
}
