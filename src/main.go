package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"crud-with-golang-and-mongodb/src/controllers"
)

func main() {
	router := mux.NewRouter()
	basePath := router.PathPrefix("/api").Subrouter()

	basePath.HandleFunc("/users", controllers.CreateProfile).Methods("POST")
	basePath.HandleFunc("/users", controllers.GetAllUsersProfile).Methods("GET")
	basePath.HandleFunc("/users", controllers.GetUserProfile).Methods("POST")
	basePath.HandleFunc("/users/{id}", controllers.UpdateUserProfile).Methods("PUT")
	basePath.HandleFunc("/users/{id}", controllers.DeleteUserProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3330", basePath))
}
