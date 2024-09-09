package controllers

import (
	"GO-BOOKAPP/pkg/models"
	"GO-BOOKAPP/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Newbook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookid := vars["bookId"]
	ID, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookdetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-TYpe", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	mybook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := mybook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookid := vars["bookId"]
	Id, err := strconv.ParseInt(bookid, 0, 0)

	if err != nil {

		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var mybook = &models.Book{}
	utils.ParseBody(r, mybook)
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	Id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookdetails, db := models.GetBookById(Id)
	if mybook.Name != "" {
		bookdetails.Name = mybook.Name
	}
	if mybook.Author != "" {
		bookdetails.Author = mybook.Author
	}
	if mybook.Publication != "" {
		bookdetails.Publication = mybook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
