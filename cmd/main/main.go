package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/florin12er/goBookstore/pkg/routes"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/",r)
    log.Fatal(http.ListenAndServe("localhost:3000",r))
}
