package routes

import (
	"github.com/Madhav8528/goBookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/v1/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/v1/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/v1/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/v1/book/{id}", controllers.DeleteBook).Methods("DELETE")

}
