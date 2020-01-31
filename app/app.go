package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
	"github.com/taka011002/go_sample_api_server/app/handler"
	"github.com/taka011002/go_sample_api_server/app/infra"
	"github.com/taka011002/go_sample_api_server/app/infra/persistence"
)

var router *mux.Router

func Run(host string) {
	router = mux.NewRouter()
	defer infra.Close()
	setRoutes()
	log.Fatal(http.ListenAndServe(host, router))
}

func setRoutes() {
	get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Pong")
	})

	userPersistence := persistence.NewUserPersistence(infra.DB)
	userService := service.NewUserService(userPersistence)
	userHandler := handler.NewUserHandler(userService)

	post("/user/login", handler.ApiHandler(userHandler.SignIn))
	post("/user", handler.ApiHandler(handler.AuthHandler(userHandler.SignUp)))
	get("/user", handler.ApiHandler(handler.AuthHandler(userHandler.GetUser)))
	put("/user", handler.ApiHandler(handler.AuthHandler(userHandler.Update)))
	delete("/user", handler.ApiHandler(handler.AuthHandler(userHandler.Delete)))
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