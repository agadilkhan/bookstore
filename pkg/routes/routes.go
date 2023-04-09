package routes

import (
	"github.com/agadilkhan/simple-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	C *controllers.Controller
}

func (router *Router) Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/books", router.C.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]}", router.C.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{title}", router.C.GetBookByTitle).Methods("GET")
	r.HandleFunc("/booksByOrder/{order}", router.C.GetBooksByOrder).Methods("GET")
	r.HandleFunc("/books", router.C.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", router.C.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}", router.C.UpdateBook).Methods("PUT")
	http.ListenAndServe(":8000", r)
}
