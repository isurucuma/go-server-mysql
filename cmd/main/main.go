package main

import (
	"github.com/gorilla/mux"
	"github.com/isurucuma/go-server-mysql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // we have routes in the routes package and we pass newly created router and regester al the routes in this router
	http.Handle("/", r)               // we pass all the routes registered router in the http default mux, router has already implemented the handler interface
	log.Fatal(http.ListenAndServe(":8080", r))
}
