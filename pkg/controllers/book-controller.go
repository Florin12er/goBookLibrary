package controllers

import (
    "fmt"
    "html/template"
    "net/http"
     "path/filepath"
    "strconv"
    "github.com/florin12er/goBookstore/pkg/models"
    "github.com/florin12er/goBookstore/pkg/utils"
    "github.com/gorilla/mux"
)

var NewBook models.Book

var templateDir, _ = filepath.Abs("../../templates")

func GetBook(w http.ResponseWriter, r *http.Request) {
    NewBooks := models.GetAllBooks()
    tmpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    tmpl.Execute(w, NewBooks)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    
    bookDetails, _ := models.GetBookById(ID)
    tmpl, err := template.ParseFiles(filepath.Join(templateDir, "edit.html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    tmpl.Execute(w, bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        CreateBook := &models.Book{}
        CreateBook.Name = r.FormValue("name")
        CreateBook.Author = r.FormValue("author")
        CreateBook.Publication = r.FormValue("publication")
        var _ = CreateBook.CreateBook()
        http.Redirect(w, r, "/book/", http.StatusSeeOther)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodDelete {
        vars := mux.Vars(r)
        bookId := vars["bookId"]
        ID, err := strconv.ParseInt(bookId, 10, 64)
        if err != nil {
            http.Error(w, "Invalid book ID", http.StatusBadRequest)
            return
        }

        // Call your delete function from models package
        err = models.DeleteBook(ID)
        if err != nil {
            http.Error(w, "Failed to delete book", http.StatusInternalServerError)
            return
        }

        // Respond with success message
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Book deleted successfully"))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var updateBook = &models.Book{}
    utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookDetails, db := models.GetBookById(ID)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }
    db.Save(&bookDetails)
    http.Redirect(w, r, fmt.Sprintf("/book/%d", bookDetails.ID), http.StatusSeeOther)
}

