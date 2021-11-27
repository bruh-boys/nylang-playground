package controllers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/labstack/echo"
)

func Share(ctx echo.Context) (err error) {
	if _, err := os.Stat("./share"); os.IsNotExist(err) {
		erro := os.Mkdir("./share", 0755)

		if erro != nil {
			log.Panicln(erro)

		}
	}

	file, _ := ioutil.TempFile("./share", "*")
	io.Copy(file, ctx.Request().Body)

	ctx.Response().Write(
		[]byte("http://" + ctx.Request().Host + "/share?code=" + file.Name()[8:]),
	)

	return
}

func Look(ctx echo.Context) (err error) {
	file := path.Clean(strings.ReplaceAll(ctx.QueryParam("code"), "..", ""))

	code, err := os.ReadFile("share/" + file)

	if err != nil {
		code = []byte("🍙 main = 🏨 () {\n\t	🎤🎶 ( \"sorry file not found XD\" ) ;\n} ;")

	}

	tpl, _ := template.ParseFiles("./public/index.html")

	tpl.Execute(
		ctx.Response(), map[string]string{"Code": string(code)},
	)

	return
}
