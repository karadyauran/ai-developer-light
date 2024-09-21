package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "./handlers"
)

func main() {
    r := mux.NewRouter()
    handlers.RegisterRoutes(r)
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}