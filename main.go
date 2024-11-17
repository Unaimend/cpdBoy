package main

import (
	"fmt"
	"net/http"
	"log"
  "github.com/Unaimend/cpdBoy/utils"
  "github.com/Unaimend/cpdBoy/server"
)

func main() {
	// Call the readTSV function
	tsvData, err := utils.ReadTSV("data/modelSeed.tsv")
	if err != nil {
		fmt.Println("Error reading TSV file:", err, )
		return
	}


  myHandler := &server.DataBaseHandler {}
  myHandler.Db = tsvData;

	// Set up routes
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		myHandler.PostMessage(w, r)
	}) 

	// Start the server
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

  

  


}

