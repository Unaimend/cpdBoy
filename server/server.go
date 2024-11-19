package server 

import (
	"net/http"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "encoding/json"
  "fmt"
  "log"
)

// Define a structure for the data you will work with
type Message struct {
	Text string 
}

type DataBaseHandler struct {
  Db sql.DB
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


   query := fmt.Sprintf(`SELECT name FROM data WHERE id = "%s"`, msg.Text)
   //fmt.Printf("\n %s \n", query)

	 // Execute the query
	 rows, err := h.Db.Query(query)
	 if err != nil {
	 	log.Fatal("Failed to execute query:", err)
	 }
	 defer rows.Close()

   for rows.Next() {
		  var name string

		  // Scan the row into variables
		  err = rows.Scan(&name)
		  if err != nil {
		  	log.Fatal("Failed to scan row:", err)
		  }
	    
      fmt.Fprintf(w, "%s,%s\n", msg.Text, name)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}



}

