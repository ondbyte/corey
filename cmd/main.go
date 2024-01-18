package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.POST("contact", h.AddContact)
	router.POST("task", h.AddTask)

	router.GET("contact/:id", h.GetContact)
	router.GET("task/:id", h.GetTask)

	router.GET("contact", h.GetAllContact)
	router.GET("task", h.GetAllTask)

	go func() {
		err = router.Run(":8000")
		if err != nil {
			panic(err)
		}
	}()
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C to exit.")
	<-signalChannel
}
