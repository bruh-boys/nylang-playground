package controller

import (
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/labstack/echo/v4"
)

func Execute(ctx echo.Context) (err error) {

	b, _ := ioutil.ReadAll(ctx.Request().Body)

	cmd := exec.Command("./nylang", ctx.Path()[1:], "-c", string(b))

	cmd.Stdout = ctx.Response()
	cmd.Stderr = ctx.Response()
	err = cmd.Run()
	go func() {
		time.Sleep(time.Minute)

		cmd.Process.Kill()
	}()
	return
}
