package controllers

import (
	"api-transaction/util"
	"encoding/json"
	"go-postgres/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "applcation/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Orign", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Fatalf("Unable to decode the equest body.  %v", err)
	}

	insertId := models.InsertTransaction(transaction)

	res := util.Response{
		ID:      insertId,
		Message: "Transaction created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	accountId := r.URL.Query().Get("accountid")
	initialDate := r.URL.Query().Get("initialdate")
	endDate := r.URL.Query().Get("enddate")

	accountIdConv, err := strconv.Atoi(accountId)
	if err != nil {
		panic(err.Error())
	}

	initialDateConv, err := time.Parse("2006-01-02", initialDate)
	if err != nil {
		panic(err.Error())
	}

	endDateConv, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		panic(err.Error())
	}

	endDateOneDayAdded := endDateConv.AddDate(0, 0, 1).Format("2006-01-02")

	transaction := models.SelectTransaction(accountIdConv, initialDateConv, endDateOneDayAdded)

	json.NewEncoder(w).Encode(transaction)
}
