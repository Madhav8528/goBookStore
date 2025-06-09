package main

import (
	"log"
	"net/http"

	"github.com/Madhav8528/goBookStore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", r))
}
