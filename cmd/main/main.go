package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Madhav8528/goBookStore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is starting on port: 9090")
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
