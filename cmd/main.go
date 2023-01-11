package main

import (
	"log"

	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/HgCl2/rock_scissors_paper/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error run http server %s", err.Error())
	}
}