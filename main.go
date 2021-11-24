package main

import (
	"github.com/bruh-boys/nylang-playground/src/router"
	"github.com/bruh-boys/nylang-playground/src/setup"
)

func main() {
	setup.SetupFiles()
	router.SetupRouter()
}
