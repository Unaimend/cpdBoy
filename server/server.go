package server 

import (
	"net/http"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "encoding/json"
  "fmt"
  "log"
  "strings"
)

type Message struct {
	Text string 
}

type DataBaseHandler struct {
  Db sql.DB
}

func QuoteAndJoin(input string) string {
	parts := strings.Split(input, ",")
	for i, part := range parts {
		parts[i] = fmt.Sprintf("\"%s\"", part)
	}
	return strings.Join(parts, ",")
}

// Handler for the /message endpoint (POST request)
func (h *DataBaseHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	// Set response header to JSON
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)
  
	var msg Message

	// Decode incoming JSON data
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Print("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

   //TODO HANDLE multiple cpds
  query := fmt.Sprintf(`SELECT name FROM data WHERE id in (%s)`, QuoteAndJoin(msg.Text))
	log.Print("query: ", query)
   
  // TODO HANDLE MULTIPLE RESULTS
	// Execute the query
	rows, err := h.Db.Query(query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer rows.Close()
  var result string

	parts := strings.Split(msg.Text, ",")
  var i = 0
  for rows.Next() {
	  var name string

	  // Scan the row into variables
	  err = rows.Scan(&name)
	  if err != nil {
	  	log.Fatal("Failed to scan row:", err)
	  }
    result += fmt.Sprintf("%s,%s\n", parts[i], name)
    i += 1
    
	  
	}
  fmt.Fprintf(w, result)

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}



}

