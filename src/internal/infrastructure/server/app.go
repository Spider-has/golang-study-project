package server

import (
	"fmt"
	"golang-web-server/src/internal/controllers"
	"golang-web-server/src/internal/infrastructure/appErrors"
	"golang-web-server/src/internal/infrastructure/config"
	"log"
	"net/http"
)


type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) IApp {
	return &App{
		cfg: cfg,
	}
}

func(a *App) Run() error {
	addr := fmt.Sprintf("%s:%s", a.cfg.Server.Host, a.cfg.Server.Port)
	fmt.Printf("Данные сервера: %s (окружение: %s)\n", addr, a.cfg.Server.Environment)

	router := controllers.NewMuxRouter();

	log.Printf("Сервер запущен на http://%s", addr)

	if err := http.ListenAndServe(addr, router.GetHandlers()); err != nil {
		return appErrors.WrapHTTPError(500, "Ошибка запуска сервера", "Не удалось запустить HTTP-сервер", err)
	}

	return nil
} 

