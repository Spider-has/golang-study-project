package main

import (
	"golang-web-server/src/internal/infrastructure/config"
	"golang-web-server/src/internal/infrastructure/server"
	"log"
)



func main() {
	cfgService := config.NewConfigService();
	cfg, err := cfgService.GetConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}