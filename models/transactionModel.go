package models

import (
	"api-transaction/dbConnection"
	"fmt"
	"log"
	"math"
	"time"
)

type Transaction struct {
	TransactionId   int     `json:"transaction_id"`
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"`
}

func InsertTransaction(transaction Transaction) int {
	db := dbConnection.CreateConnection()

	sqlStatement := `INSERT INTO "Transactions" ("Account_ID", "OperationType_ID", "Amount", "EventDate" ) 
	VALUES ($1, $2, $3, current_timestamp) RETURNING "Transaction_ID"`

	var transactionId int
	absoluteAmount := math.Abs(transaction.Amount)

	err := db.QueryRow(sqlStatement, transaction.AccountId, transaction.OperationTypeId, absoluteAmount).Scan(&transactionId)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer db.Close()

	fmt.Println("Transaction inserted %v", transactionId)

	return transactionId
}

func SelectTransaction(accountId int, initialDate time.Time, endDate string) []Transaction {
	db := dbConnection.CreateConnection()

	sqlStatement := `SELECT t."Transaction_ID", t."Account_ID", t."OperationType_ID", (t."Amount" * o."Sign") as "Amount", t."EventDate"
	FROM "Transactions" t
	JOIN "OperationsTypes" o ON t."OperationType_ID" = o."OperationType_ID"
	WHERE "Account_ID" = $1 
	AND "EventDate" >= $2 
	AND "EventDate" < $3
	ORDER BY "EventDate"`

	rows, err := db.Query(sqlStatement, accountId, initialDate, endDate)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	item := Transaction{}
	transactions := []Transaction{}

	for rows.Next() {
		var date time.Time
		var transactionId, accountId, operationTypeId int
		var amount float64

		err := rows.Scan(&transactionId, &accountId, &operationTypeId, &amount, &date)
		if err != nil {
			panic(err.Error())
		}

		item.TransactionId = transactionId
		item.AccountId = accountId
		item.OperationTypeId = operationTypeId
		item.Amount = amount
		item.EventDate = date.Local().Format("2006-01-02  15:04:05")

		transactions = append(transactions, item)
	}

	defer db.Close()

	return transactions
}
