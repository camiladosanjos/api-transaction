package router

import (
	"api-transaction/controllers"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST", "OPTIONS")
	router.HandleFunc("/accounts/{accountId}", controllers.GetAccount).Methods("GET", "OPTIONS")
	router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST", "OPTIONS")
	router.HandleFunc("/transactions", controllers.GetTransaction).Methods("GET", "OPTIONS")

	return router
}
