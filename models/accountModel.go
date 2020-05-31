package models

import (
	"api-transaction/dbConnection"
	"fmt"
	"log"
)

type Account struct {
	AccountId int    `json:"account_id"`
	Document  string `json:"document_number"`
}

func InsertAccount(account Account) int {
	db := dbConnection.CreateConnection()

	sqlStatement := `INSERT INTO "Accounts" ("Document_Number") VALUES ($1) RETURNING "Account_ID"`

	var accountId int

	err := db.QueryRow(sqlStatement, account.Document).Scan(&accountId)
	if err != nil {
		log.Fatalf("Unable to execute the query.", err)
	}

	defer db.Close()
	fmt.Println("Account inserted %v", accountId)

	return accountId
}

func SelectAccount(accountId int) (Account, error) {
	db := dbConnection.CreateConnection()

	var account Account

	sqlStatement := `SELECT * FROM "Accounts" WHERE "Account_ID"=$1`

	row := db.QueryRow(sqlStatement, accountId)

	err := row.Scan(&account.AccountId, &account.Document)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return account, err
}
