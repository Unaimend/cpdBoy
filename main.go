package main

// Use this intead of custom tsv handler
// https://earthly.dev/blog/golang-csv-files/#:~:text=data%20in%20Golang.-,Using%20Encoding%2Fcsv%20To%20Work%20with%20CSV%20Files,read%20and%20write%20CSV%20data.

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Unaimend/cpdBoy/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

  // Creating log file
  logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil {
      log.Fatal(err)
  }
  defer logFile.Close()

  //Init db and slog
  myHandler := &server.DataBaseHandler {}
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  slog.LogAttrs(context.Background(), slog.LevelInfo, "Starting the server")
  slog.SetDefault(logger)
  db, err := sql.Open("sqlite3", "./my_database.db")
  myHandler.Db = *db


  if err != nil {
    log.Fatal("Failed to open the database:", err)
  }
  slog.Info("Opened the database successfully")
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
  slog.Info(fmt.Sprintf("Server is running on http://localhost%s\n", port))
  if err := http.ListenAndServe(port, nil); err != nil {
  	log.Fatalf("Error starting server: %v", err)
  }
}

