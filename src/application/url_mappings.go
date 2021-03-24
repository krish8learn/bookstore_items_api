package application

import (
	"net/http"

	"github.com/krish8learn/bookstore_items_api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	//router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	//router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	//router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
