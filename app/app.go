package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
	"github.com/taka011002/go_sample_api_server/app/handler"
	"github.com/taka011002/go_sample_api_server/app/infra"
	"github.com/taka011002/go_sample_api_server/app/infra/persistence"
	"log"
	"net/http"
)

var router *mux.Router

func Run(host string) {
	router = mux.NewRouter()
	setRoutes()
	log.Fatal(http.ListenAndServe(host, router))
}

func setRoutes() {
	get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ping")
	})

	userPersistence := persistence.NewUserPersistence(infra.DB)
	userService := service.NewUserService(userPersistence)
	userHandler := handler.NewUserHandler(userService)

	get("/user/{username}", userHandler.GetUser)
	post("/user", userHandler.CreateUser)
	get("/user", handleRequest(userHandler.GetUser))
	post("/user", handleRequest(userHandler.UpdateUser))
}

// Get wraps the router for GET method
func get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("DELETE")
}

func handleRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}