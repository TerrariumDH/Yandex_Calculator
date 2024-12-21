package main

import (
	"github.com/TerrariumDH/Yandex_Calculator/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}
