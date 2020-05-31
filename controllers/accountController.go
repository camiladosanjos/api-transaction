package controllers

import (
	"api-transaction/util"
	"encoding/json"
	"go-postgres/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Fatalf("Unable to decde the request body.  %v", err)
	}

	insertID := models.InsertAccount(account)

	res := util.Response{
		ID:      insertID,
		Message: "Account created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	accountId, err := strconv.Atoi(params["accountId"])
	if err != nil {
		log.Fatalf("Unable t convert the string into int.  %v", err)
	}

	account, err := models.SelectAccount(int(accountId))
	if err != nil {
		log.Fatalf("Unable to get account. %v", err)
	}

	json.NewEncoder(w).Encode(account)
}
