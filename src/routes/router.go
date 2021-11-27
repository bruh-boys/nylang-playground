package router

import (
	"log"
	"os"

	"github.com/bruh-boys/nylang-playground/src/controllers"
	"github.com/labstack/echo"
)

func SetupRouter() {
	app := echo.New()

	app.Static("/codemirror", "./codemirror")
	app.Static("/public", "public")

	app.POST("/run", controllers.RunCode)

	app.POST("/lexer", controllers.RunCode)

	app.POST("/parser", controllers.RunCode)

	app.POST("/share", controllers.Share)
	app.GET("/share", controllers.Look)

	app.GET("/", func(c echo.Context) error {
		return c.File("./public/index.html")

	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"

	}

	log.Panicln(app.Start(":" + port))
}
