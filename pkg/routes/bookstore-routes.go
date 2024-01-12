package routes

import (
	"github.com/gorilla/mux"
	"github.com/Mikael88/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	const rute string = "/book/{bookId}" 
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc(rute, controllers.GetBookById).Methods("GET")
	router.HandleFunc(rute, controllers.UpdateBook).Methods("PUT")
	router.HandleFunc(rute, controllers.DeleteBook).Methods("DELETE")
}