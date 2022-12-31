package routes

import (
	"github.com/gorilla/mux"
	 "github.com/gbubemi22/go-bookstore/pkg/controllers"
)


var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
     router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
     router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}