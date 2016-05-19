package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)


const (
	DB_USER     = "postgres"
	DB_PASSWORD = "pass_1"
	DB_NAME     = "postgres"
)


func main() {
	dbconnect := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbconnect)
	checkErr(err)
	defer db.Close()


}


func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}