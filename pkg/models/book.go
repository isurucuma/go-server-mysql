package models

import (
	"github.com/isurucuma/go-server-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book this is the model for the book table
// we have used the struct tags to give metadata to the fields
// users of this struct can use these tags to get the metadata and use them for validations
type Book struct {
	gorm.Model // this is used to add the id, created_at, updated_at, deleted_at fields to the model
	// in go we can embed a struct to another struct then the fields of the embedded struct
	// will be added to the parent struct
	Name        string `gorm:"" json:"name"` // gorm:"" is used to tell gorm that this field is not nullable
	Author      string `json:"author"`       //
	Publication string `json:"publication"`
}

/**
When a Go program starts, the init functions of the imported packages are executed first in a depth-first order,
meaning that the init functions of the directly imported packages are called before the init functions of their dependencies.
Once all the init functions of the imported packages have been executed, the main function of the main package is called
to start the program execution
*/

// init this is used to initialize the database connection
// this is a special function in go which gets called before the main function
// which is a special function that is automatically called by the Go runtime before the main function in the same package.
// The purpose of init functions is to perform package-level initialization tasks,
// such as setting up variables, registering components, or initializing state.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // this is used to migrate the model to the database
}

// CreateBook this method is used to create a new book in the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // this is used to tell gorm that the record doesn't exist in the database, then gorm will create a new record with the data given
	db.Create(&b)   // this is used to create a new record in the database
	return b
}

// GetAllBooks this method is used to get all the books from the database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) // finds and put the result in the Books variable
	return Books
}

// GetBookById this method is used to get a book by its id
func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// DeleteBook this method is used to delete a book by its id
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
