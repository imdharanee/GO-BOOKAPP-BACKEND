package main

import (
	"log"
	"net/http"

	"GO-BOOKAPP/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))

}
