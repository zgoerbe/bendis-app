package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/zgoerbe/bendis"
)

type application struct {
	App        *bendis.Bendis
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	b := initApplication()
	go b.listenForShutdown()
	err := b.App.ListenAndServer()
	b.App.ErrorLog.Println(err)
}

func (a *application) shutDown() {
	// put any clean up tasks here

	// block until wait group is empty
	a.wg.Wait()
}

func (a *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	a.App.InfoLog.Println("Received signal", s.String())
	a.shutDown()

	os.Exit(0)
}
