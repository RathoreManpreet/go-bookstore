package routes

import (
	"github.com/RathoreManpreet/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisteredBookStoreRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
