package models

import (
	"GO-BOOKAPP/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var mybook Book
	db := db.Where("ID=?", Id).Find(&mybook)
	return &mybook, db

}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
