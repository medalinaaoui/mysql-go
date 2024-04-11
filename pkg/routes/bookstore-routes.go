package routes

import (
	"github.com/gorilla/mux"
	"github.com/test/mysql-go/pkg/controllers"
)



func RegisterBookStoreRoutes(router *mux.Router){
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateSingleBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteSingleBook).Methods("DELETE")
} 