package main

import (
	router "github.com/bruh-boys/nylang-playground/src/routes"
	"github.com/bruh-boys/nylang-playground/src/setup"
)

func main() {
	setup.SetupFiles()
	router.SetupRouter()
}
