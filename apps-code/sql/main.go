package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Video struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/addVideo", addVideo)
	http.HandleFunc("/getVideos", getVideos)
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
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	url := req.URL.Query().Get("url")
	query := fmt.Sprintf(
		"INSERT INTO videos VALUES ('%s', '%s', '%s')",
		id,
		name,
		url,
	)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "All is good.")
	defer insert.Close()
}

func getVideos(w http.ResponseWriter, req *http.Request) {
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
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := "SELECT id, name, url FROM videos"
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var video Video
		err = results.Scan(&video.ID, &video.Name, &video.URL)
		if err != nil {
			panic(err.Error())
		}
		output := fmt.Sprintf(
			"ID: %s\nName: %s\nURL: %s\n----------\n",
			video.ID,
			video.Name,
			video.URL,
		)
		log.Printf(output)
		fmt.Fprintf(w, output)
	}

	defer results.Close()
}
