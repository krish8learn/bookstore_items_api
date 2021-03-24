package main

import (
	"github.com/gorilla/mux"
	"github.com/krish8learn/bookstore_items_api/src/application"
)

var (
	router = mux.NewRouter()
)

func main() {
	application.StartApplication()
}
