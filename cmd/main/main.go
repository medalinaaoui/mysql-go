package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/medali/test/pkg/routes"
)

	
func main() {


	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8001", r); err == nil {
		log.Fatal(err)
	}

}