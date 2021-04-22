package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/dbijay/simple-gorest-crud/pkg/models"
// 	"github.com/dbijay/simple-gorest-crud/pkg/utils"
// 	"github.com/gorilla/mux"
// )

// var NoteBook models.Book

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	b := CreateBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	// newBooks := models.GetAllBooks()
// 	AllBooks := &models.Book{}
// 	res, _ := json.Marshal(AllBooks.GetAllBooks())
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	// bookDetails, _ := models.GetBookById(ID)
// 	OneBook := &models.Book{}
// 	bookDetails, _ := OneBook.GetBookById(ID)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	// bookDetails, db := models.GetBookById(ID)
// 	OneBook := &models.Book{}
// 	bookDetails, db := OneBook.GetBookById(ID)
// 	if updateBook.Name != "" {
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != "" {
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != "" {
// 		bookDetails.Publication = updateBook.Publication
// 	}
// 	db.Save(&bookDetails)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	// book := models.DeleteBook(ID)
// 	DelBook := &models.Book{}
// 	res, _ := json.Marshal(DelBook.DeleteBook(ID))
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
