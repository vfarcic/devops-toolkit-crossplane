package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/addVideo", addVideo)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Println("Starting the server on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func root(w http.ResponseWriter, req *http.Request) {
}

func addVideo(w http.ResponseWriter, req *http.Request) {
	log.Println("Adding a video...")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	dbPort := os.Getenv("DB_PORT")
	dbUri := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tls=false",
		dbUser,
		dbPass,
		dbEndpoint,
		dbPort,
		dbName,
	)
	log.Println(dbUri)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	id := req.URL.Query().Get("id")
	title := req.URL.Query().Get("title")
	url := req.URL.Query().Get("url")
	query := fmt.Sprintf(
		"INSERT INTO videos VALUES ('%s', '%s', '%s')",
		id,
		title,
		url,
	)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
