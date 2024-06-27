package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
    if !checkAuth(r) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Invalid user ID"}`))
        return
    }

    user, exists := users[userID]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "User not found"}`))
        return
    }

    json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    if !checkAuth(r) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    var userList []User
    for _, user := range users {
        userList = append(userList, user)
    }

    json.NewEncoder(w).Encode(userList)
}
