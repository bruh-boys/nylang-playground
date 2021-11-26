package controller

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

func Share(ctx echo.Context) (err error) {

	f, _ := ioutil.TempFile("./share", "*")

	io.Copy(f, ctx.Request().Body)
	log.Println(f.Name())
	ctx.Response().Write([]byte("http://" + ctx.Request().Host + "/share?code=" + f.Name()[8:]))
	return

}

func Look(ctx echo.Context) (err error) {
	file := path.Clean(strings.ReplaceAll(ctx.QueryParam("code"), "..", ""))
	b, err := os.ReadFile("share/" + file)
	log.Println("share/" + file)
	if err != nil {
		b = []byte("🍙 main = 🏨 () {\n\t	🎤🎶 ( \"sorry file not found XD\" ) ;\n} ;")
	}

	t, _ := template.ParseFiles("./template/index.html")
	t.Execute(ctx.Response(), map[string]string{"Code": string(b)})

	return
}
