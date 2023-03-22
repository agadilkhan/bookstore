package main

import (
	"github.com/agadilkhan/simple-rest-api/pkg/controllers"
	"github.com/agadilkhan/simple-rest-api/pkg/routes"
	"github.com/agadilkhan/simple-rest-api/pkg/store"
)

func main() {
	var store store.Store
	store.Connection()
	var controller controllers.Controller = controllers.Controller{DB: &store}
	r := routes.Router{C: &controller}
	r.Routes()
}
