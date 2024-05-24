package main

import (
	"go-init/config"
	categorycontroller "go-init/controllers/categoryController"
	homecontroller "go-init/controllers/homeController"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// Home
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8000", nil)
}
