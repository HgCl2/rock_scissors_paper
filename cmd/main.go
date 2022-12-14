package main

import (
	"log"

	app "github.com/HgCl2/rock_scissors_paper"
)

func main() {
	srv := new(app.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error run http server %s", err.Error())
	}
}
