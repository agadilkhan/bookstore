package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookByID).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("books/{id}", DeleteBook).Methods("DELETE")

	return r
}