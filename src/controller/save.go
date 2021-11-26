package controller

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Share(ctx echo.Context) (err error) {

	f, _ := ioutil.TempFile("./share", "*")

	io.Copy(f, ctx.Request().Body)

	ctx.Redirect(http.StatusTemporaryRedirect, "/share?id="+f.Name())
	return

}

func Look(ctx echo.Context) (err error) {

	return
}
