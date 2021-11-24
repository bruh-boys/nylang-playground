package router

import (
	"os"

	"github.com/bruh-boys/nylang-playground/src/controller"
	"github.com/labstack/echo/v4"
)

func SetupRouter() {
	app := echo.New()
	app.POST("/run", controller.Execute)
	app.POST("/lexer", controller.Execute)
	app.POST("/parser", controller.Execute)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Logger.Fatal(app.Start(":" + port))
}
