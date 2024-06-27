package main

import (
    "net/http"
    "strings"
)

func checkAuth(r *http.Request) bool {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return false
    }

    authParts := strings.SplitN(authHeader, " ", 2)
    if len(authParts) != 2 || authParts[0] != "Basic" {
        return false
    }

    auth := strings.SplitN(authParts[1], ":", 2)
    if len(auth) != 2 {
        return false
    }

    username, password := auth[0], auth[1]
    return username == "988d4f1313f881e5ac6bfdfc7f54244aab" && password == "905a12r3a0c"
}
