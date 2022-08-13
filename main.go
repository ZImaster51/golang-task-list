package main

import (
	"fmt"
	"golang-task-list/config"
	"golang-task-list/controllers"
	"golang-task-list/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.ConnectionDatabase()

	if err != nil {
		panic("Program error bos")
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteControllers := &controllers.NoteControllers{}

	router := httprouter.New()

	router.GET("/", noteControllers.Index)
	router.GET("/create", noteControllers.Create)
	router.POST("/create", noteControllers.Store)
	router.GET("/edit/:id", noteControllers.Edit)
	router.POST("/edit/:id", noteControllers.Update)
	router.POST("/done/:id", noteControllers.Done)
	router.POST("/delete/:id", noteControllers.Delete)

	port := ":2368"
	fmt.Println("Aplikasi jalan di http://localhost:2368")

	// fmt.Println("aman boss")
	log.Fatal(http.ListenAndServe(port, router))

}
