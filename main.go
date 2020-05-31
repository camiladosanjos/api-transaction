package main

import (
	"api-transaction/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
