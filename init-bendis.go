package main

import (
	"log"
	"myapp_pt2/data"
	"myapp_pt2/handlers"
	"myapp_pt2/middleware"
	"os"

	"github.com/zgoerbe/bendis"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init bendis
	bend := &bendis.Bendis{}
	err = bend.New(path)
	if err != nil {
		log.Fatal(err)
	}

	bend.AppName = "myapp_pt2"
	
	myMiddleware := &middleware.Middleware{
		App:    bend,
	}

	//bend.InfoLog.Println("Debug is set to", bend.Debug)
	myHandlers := &handlers.Handlers{
		App: bend,
	}

	app := &application{
		App: bend,
		Handlers: myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
