package main

import (
	"log"

	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/HgCl2/rock_scissors_paper/pkg/handler"
	"github.com/HgCl2/rock_scissors_paper/pkg/repository"
	"github.com/HgCl2/rock_scissors_paper/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error run http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
