package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")
	// Change to env vars
	db, err := sql.Open("mysql", "root:yVUP9OXMC7GqCksrE872UNffftw@tcp(34.73.107.106:3306)/dot?tls=false")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("111")
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	fmt.Println("222")
	if err != nil {
		fmt.Println("333")
		panic(err.Error())
	}
	fmt.Println("444")
	defer insert.Close()
}
