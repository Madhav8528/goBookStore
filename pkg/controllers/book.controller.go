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
	var book *models.Book
	utils.ParseBody(r, book)
	b := models.CreateBook(book)
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

	//delete the entry
	models.DeleteBook(int64(id))

	//create updated entry
	b := models.CreateBook(book)
	res, _ := json.Marshal(&b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "pkglication/json")

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 0, 0)

	b := models.DeleteBook(int64(id))
	res, _ := json.Marshal(&b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
