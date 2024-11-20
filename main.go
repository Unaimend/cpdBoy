package main

// Use this intead of custom tsv handler
// https://earthly.dev/blog/golang-csv-files/#:~:text=data%20in%20Golang.-,Using%20Encoding%2Fcsv%20To%20Work%20with%20CSV%20Files,read%20and%20write%20CSV%20data.

import (
	"net/http"
	"log"
  "github.com/Unaimend/cpdBoy/server"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

func main() {
	// Call the readTSV function


  myHandler := &server.DataBaseHandler {}
  db, err := sql.Open("sqlite3", "./my_database.db")
  myHandler.Db = *db
  
  if err != nil {
    log.Fatal("Failed to open the database:", err)
  }
  defer db.Close()
  
  // Check the connection
  if err := db.Ping(); err != nil {
    log.Fatal("Failed to connect to the database:", err)
  }


	// Set up routes
  http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
  	myHandler.PostMessage(w, r)
  }) 
  
  // Start the server
  port := ":3000"
  log.Printf("Server is running on http://localhost%s\n", port)
  if err := http.ListenAndServe(port, nil); err != nil {
  	log.Fatalf("Error starting server: %v", err)
  }
}

