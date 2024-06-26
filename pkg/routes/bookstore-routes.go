package routes

import (
    "github.com/gorilla/mux"
    "github.com/florin12er/goBookstore/pkg/controllers"
)

// RegisterBookStoreRoutes registers all routes related to bookstore operations
func RegisterBookStoreRoutes(router *mux.Router) {
    // Endpoint for creating a new book (POST /book/)
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

    // Endpoint for retrieving all books (GET /book/)
    router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

    // Endpoint for retrieving a specific book by ID (GET /book/{bookId})
    router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

    // Endpoint for updating a specific book by ID (PUT /book/{bookId})
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

    // Endpoint for deleting a specific book by ID (DELETE /book/{bookId})
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

