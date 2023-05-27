package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/isurucuma/go-server-mysql/pkg/models"
	"github.com/isurucuma/go-server-mysql/pkg/utils"
	"net/http"
	"strconv"
)

// GetBooks this function is used to get all the books from the database
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // we call the GetAllBooks method to get all the books
	res, _ := json.Marshal(newBooks) // we convert the response to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // writer writes the response to the response writes as the method gets returned
}

// GetBook this function is used to get a single book from the database
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // this is used to get the path params in the request, this returns a key value map
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0) // we convert the string to int64, the 0,0 is used to tell the parse method that the string is base 10 and 64 bit integer
	if err != nil {
		fmt.Println("error while parsing the bookId")
	}
	bookDetails, _ := models.GetBookById(Id) // get the book using the ID passed by the user
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook this function is used to create a new book in the database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createdBook := &models.Book{}   // we create a new book object and new data added to it
	utils.ParseBody(r, createdBook) // here we parse the request body and map it to the createdBook object, after this createdBook will have all the data that the user has sent in the request
	b := createdBook.CreateBook()   // we call the CreateBook method with the createdBook object as the receiver
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// UpdateBook this function is used to update a book in the database
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// get the available book from the database using the Id passed by the user
	bookDetails, db := models.GetBookById(Id)
	// if user passed the name then update the name
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Name != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Name != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(bookDetails) // save the updated book
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// DeleteBook this function is used to delete a book from the database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing the bookId")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
