package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/onbyte/corey"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(corey.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = corey.MigrateModels(db)
	if err != nil {
		log.Fatal(err)
	}
	r := corey.NewRepo(db)
	s := corey.NewService(r)
	h := corey.NewHandler(s)

	//chi
	chiRouter := chi.NewRouter()
	chiRouter.Post("/contact", h.AddContact)
	chiRouter.Post("/task", h.AddTask)

	chiRouter.Get("/contact/{id}", h.GetContact)
	chiRouter.Get("/task/{id}", h.GetTask)

	chiRouter.Get("/contact", h.GetAllContact)
	chiRouter.Get("/task", h.GetAllTask)

	go func() {
		err = http.ListenAndServe(":8000", chiRouter)
		if err != nil {
			panic(err)
		}
	}()
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C to exit.")
	<-signalChannel
}
