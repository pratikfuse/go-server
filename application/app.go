package application

import (
	"fiberapp/application/controllers"
	"fiberapp/application/response"
	_ "fiberapp/database"
	"os"

	"github.com/gofiber/fiber"
)

type App struct {
	Server *fiber.App
}

func (app *App) Start() error {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "5000"
	}
	err := app.Server.Listen(appPort)
	return err
}

func Initialize() *App {
	app := App{Server: fiber.New()}
	router := controllers.Router{App: app.Server}
	app.Server.Get("/health", func(ctx *fiber.Ctx) {
		ctx.JSON(response.Json{"status": "ok"})
	})
	router.InitRoutes()
	return &app
}
