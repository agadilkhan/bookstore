package controllers

import (
	"encoding/json"
	"github.com/agadilkhan/simple-rest-api/pkg/models"
	"github.com/agadilkhan/simple-rest-api/pkg/store"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {
	DB *store.Store
}

func (c *Controller) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.DB.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	err = c.DB.CreateBook(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func (c *Controller) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := c.DB.GetBookByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := c.DB.DeleteBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func (c *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	err = c.DB.UpdateBook(id, &updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func (c *Controller) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, _ := vars["title"]
	book, err := c.DB.SearchBook(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
