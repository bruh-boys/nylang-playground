package setup

import (
	"io"
	"net/http"
	"os"
	"runtime"
)

var (
	links = map[string]string{
		"darwin": "https://github.com/bruh-boys/nylang/releases/download/nylang/darwin",
		"linux":  "https://github.com/bruh-boys/nylang/releases/download/nylang/linux",
	}
)

func downloadFile() {
	link, e := links[runtime.GOOS]
	if !e {
		panic("please dont use windows, we hate windows in this organization")
	}
	r, _ := http.Get(link) //if it has an error i dont want to continue , thats the reason why im not handling it
	f, _ := os.OpenFile("nylang", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	io.Copy(f, r.Body)
}
