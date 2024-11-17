package server 

import (
	"encoding/json"
	"fmt"
	"net/http"
  "github.com/Unaimend/cpdBoy/utils"
)

// Define a structure for the data you will work with
type Message struct {
	Text string 
}

type DataBaseHandler struct {
  Db utils.DataBase 
}

// Handler for the /message endpoint (POST request)
func (h *DataBaseHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	// Set response header to JSON
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)

	var msg Message

	// Decode incoming JSON data
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


  row := utils.FilterBy(h.Db, "id", msg.Text)
  cell := row[0]["name"]

	// Respond with the received message
	fmt.Fprintf(w, "%s,%s\n", msg.Text, cell)

}

