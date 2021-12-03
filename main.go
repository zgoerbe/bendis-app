package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/zgoerbe/bendis"
)

type application struct {
	App        *bendis.Bendis
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	b := initApplication()
	b.App.ListenAndServer()

}
