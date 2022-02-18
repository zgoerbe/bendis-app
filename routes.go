package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/zgoerbe/bendis"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.get("/", a.Handlers.Home)
	//a.get("/tester", a.Handlers.Clicker)

	//a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	//a.App.Routes.Post("/users/login", a.Handlers.PostUserLogin)
	//a.get("/users/login", a.Handlers.UserLogin)
	//a.post("/users/login", a.Handlers.PostUserLogin)
	//a.get("/users/logout", a.Handlers.Logout)
	//
	//a.get("/auth/{provider}", a.Handlers.SocialLogin)
	//a.get("/auth/{provider}/callback", a.Handlers.SocialMediaCallback)
	//
	//a.get("/upload", a.Handlers.BendisUpload)
	//a.post("/upload", a.Handlers.PostBendisUpload)
	//
	//a.get("/list-fs", a.Handlers.ListFS)
	//
	//a.get("/files/upload", a.Handlers.UploadToFS)
	//a.post("/files/upload", a.Handlers.PostUploadToFS)
	//
	//a.get("/delete-from-fs", a.Handlers.DeleteFromFS)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from bendis
	a.App.Routes.Mount("/bendis", bendis.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
