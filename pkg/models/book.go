package models

// import (
// 	"github.com/dbijay/simple-gorest-crud/pkg/config"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// type Book struct {
// 	gorm.Model
// 	// Id string `json:"id"`
// 	// Name        string `gorm:""json:"name"`
// 	Name        string `json:"name"`
// 	Author      string `json:"author"`
// 	Publication string `json:"publication"`
// }

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&Book{})
// }

// func (b *Book) CreateBook() *Book {
// 	// db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

// func (b *Book) GetAllBooks() []Book {
// 	var Books []Book
// 	db.Find(&Books)
// 	return Books
// }

// func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID = ?", Id).Find(&getBook)
// 	return &getBook, db
// }

// func (b *Book) DeleteBook(ID int64) string {
// 	// var book := &Book
// 	// db.Where("ID = ?", ID).Delete(&Book{}, ID)
// 	db.Unscoped().Delete(&Book{}, ID)
// 	return "Deleted Sucessfully"

// 	// AllBooks := &Book{}
// 	// res := AllBooks.GetAllBooks()
// 	// return res
// }
