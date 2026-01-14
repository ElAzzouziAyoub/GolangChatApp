package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
)


type Message struct{
	From string		`json:"body"`
	To string		`json:"body"`
	Body string `json:"body"`

}

var inboxes = make(map[string][]Message)
var constString string = "postgres://ayoub:secret@localhost:5432/testdb?sslmode=disable"
var db *sql.DB 
var err error


func handle(w http.ResponseWriter,r *http.Request){
	fmt.Println("Connection established")

	switch r.Method {
	case http.MethodPost:
				var msg Message
        if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        // Insert into Postgres
        _, err := db.Exec(
            "INSERT INTO messages (sender, receiver, body) VALUES ($1,$2,$3)",
            msg.From, msg.To, msg.Body,
        )
        if err != nil {
            log.Println("Insert failed:", err)
            http.Error(w, "DB error", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Message stored successfully"))
	

	case http.MethodGet:
    receiver := r.URL.Query().Get("receiver")

    rows, err := db.Query("SELECT sender, body FROM messages WHERE receiver=$1 ORDER BY id", receiver)
    if err != nil {
        http.Error(w, "DB error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var messages []Message
    for rows.Next() {
        var m Message
        if err := rows.Scan(&m.From, &m.Body); err != nil {
            continue
        }
        messages = append(messages, m)
    }

    json.NewEncoder(w).Encode(messages)

	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main(){

	db ,err = sql.Open("postgres",constString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/",handle)
	http.ListenAndServe(":8080",nil)

}
