package controllers

import (
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/labstack/echo"
)

func RunCode(ctx echo.Context) (err error) {
	code, _ := ioutil.ReadAll(ctx.Request().Body)

	cmd := exec.Command("./nylang", ctx.Path()[1:], "-c", string(code))

	go func() {
		time.Sleep(time.Second * 10)
		cmd.Process.Kill()
	}()

	cmd.Stdout = ctx.Response()
	cmd.Stderr = ctx.Response()

	cmd.Run()

	return
}
