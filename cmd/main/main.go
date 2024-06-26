package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/florin12er/goBookstore/pkg/routes"
)
func MethodOverride(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            if r.FormValue("_method") == "DELETE" {
                r.Method = http.MethodDelete
            }
        }
        next.ServeHTTP(w, r)
    })
}
func main() {
    r := mux.NewRouter()
    r.Use(MethodOverride)
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/",r)
    addr := "localhost:3000"
    log.Printf("Starting server on %s", addr)
    log.Fatal(http.ListenAndServe(addr,r))
}
