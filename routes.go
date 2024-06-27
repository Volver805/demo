package routes

import (
    "github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {
    r.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/api/v1/users/{id}", handlers.GetUser).Methods("GET")
}
