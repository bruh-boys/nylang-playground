package controller

import (
	"io"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func Share(ctx echo.Context) (err error) {

	f, _ := ioutil.TempFile("./share", "*.nyl")

	io.Copy(f, ctx.Request().Body)
	return

}

func Look(ctx echo.Context) (err error) {

	return
}
