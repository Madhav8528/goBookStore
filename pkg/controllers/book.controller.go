package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Madhav8528/goBookStore/pkg/models"
	"github.com/Madhav8528/goBookStore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")

	books := models.GetAllBooks()
	res, _ := json.Marshal(&books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while converting book id to int", err)
	}

	book, _ := models.GetBook(id)
	res, _ := json.Marshal(&book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")
	var book *models.Book

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 0, 0)
	utils.ParseBody(r, &book)

	//find the entry
	b, db := models.GetBook(int64(id))

	//update the entry
	if book.Author != "" {
		b.Author = book.Author
	}
	if book.Name != "" {
		b.Name = book.Name
	}
	if book.Publication != "" {
		b.Publication = book.Publication
	}
	db.Save(&b)
	res, _ := json.Marshal(&b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while converting book id to int", err)
	}

	b := models.DeleteBook(int64(id))
	res, _ := json.Marshal(&b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
