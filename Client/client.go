package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq" // Postgres driver
	"net/http"
	"encoding/json"
	"bytes"
)


type Message struct{
	From string		`json:"from"`
	To string		`json:"to"`
	Body string `json:"body"`

}

var db *sql.DB
var err error


var from string
var inbox []string
var constString string = "postgres://ayoub:secret@localhost:5432/testdb?sslmode=disable"

func SendMessage(){
	
			var to string 
			var body string 

			fmt.Print("\n Who do you want to send the message to :")
			fmt.Scan(&to)
			fmt.Print("\n--->")
			fmt.Scan(&body)

			msg := Message{
				From:from,
				To:to,
				Body:body,
			}


			data, err := json.Marshal(msg)
    	if err != nil {
        log.Fatal("JSON marshal error:", err)
    	}

    	resp, err := http.Post(
        "http://localhost:8080/",
        "application/json",
        bytes.NewBuffer(data),
    	)
    	if err != nil {
        log.Fatal("HTTP POST error:", err)
    	}
    	defer resp.Body.Close()

			
}


func CheckInbox(username string) {
    // Build the URL with query parameter
    url := fmt.Sprintf("http://localhost:8080/?receiver=%s", username)

    // Send GET request to server
    resp, err := http.Get(url)
    if err != nil {
        log.Println("HTTP GET failed:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Println("Server returned:", resp.Status)
        return
    }

    // Decode the JSON array returned by server
    var messages []Message
    if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
        log.Println("Failed to decode JSON:", err)
        return
    }

    // Display messages
    if len(messages) == 0 {
        fmt.Println("Inbox is empty.")
        return
    }

    fmt.Println("\n--- Inbox ---")
    for _, m := range messages {
        fmt.Printf("From %s: %s\n", m.From, m.Body)
    }
    fmt.Println("------------")
}



func main(){
	db ,err = sql.Open("postgres",constString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Print("Enter your username :")
	fmt.Scanln(&from)
	fmt.Printf("\nWelcome %s To the GolangChatApp \n")


	for {
		var choice int 
		
		fmt.Println("---------------------")
		fmt.Println("Choose an Option :")
		fmt.Println("1-Send a message ")
		fmt.Println("2- Check Inbox")
		fmt.Println("9-Quit ")
		fmt.Print("-----> ")

		fmt.Scan(&choice)

		if choice == 9 {
			log.Fatal("Quitted !")
		}
		if choice == 1 {
			SendMessage()
		}
		if choice == 2 {
			CheckInbox(from)
		}



	}

}

