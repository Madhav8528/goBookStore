package routes

import (
	"github.com/Madhav8528/goBookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}
