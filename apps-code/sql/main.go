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
	log.Println("Not yet implemented...")
}

func addVideo(w http.ResponseWriter, req *http.Request) {
	log.Println("Adding a video...")
	db, err := sql.Open("mysql", "root:OSAldwySWcXNxxUCCMlXlXUjU05@tcp(35.229.106.29:3306)/sql-demo?tls=false")
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
