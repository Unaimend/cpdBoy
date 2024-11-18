package main

// Use this intead of custom tsv handler
// https://earthly.dev/blog/golang-csv-files/#:~:text=data%20in%20Golang.-,Using%20Encoding%2Fcsv%20To%20Work%20with%20CSV%20Files,read%20and%20write%20CSV%20data.

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
	port := ":3000"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

  

  


}

