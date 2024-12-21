package main

import (
	"github.com/Ilya-goose/Calc_http_go/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}
