package routes

import (
	"github.com/gorilla/mux"
	"github.com/isurucuma/go-server-mysql/pkg/controllers"
)

// RegisterBookStoreRoutes this function gets a router pointer as input and register all the routes in it
func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
