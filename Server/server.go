package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"database/sql"
)


type Message struct{
	From string		`json:"body"`
	To string		`json:"body"`
	Body string `json:"body"`

}

var inboxes = make(map[string][]Message)
var mu sync.Mutex // protects inboxes
var constString string = "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"
var db *sql.DB 
var err error


func handle(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Wassup"))
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

	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main(){

	db ,err := sql.Open("postgres",constString)
		defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/",handle)
	http.ListenAndServe(":8080",nil)

}
