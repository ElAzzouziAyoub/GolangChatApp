package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq" // Postgres driver
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

			something, err := db.Exec(
    		"INSERT INTO messages (sender, receiver, body) VALUES ($1, $2, $3)",
    		msg.From,
    		msg.To,
    		msg.Body,
			)
			if err != nil {
    		log.Println("insert failed:", err)
			}
			if something == nil {
				fmt.Print("nothing literally")
			}
}



func CheckInbox() {
    rows, err := db.Query(
        "SELECT sender, body FROM messages WHERE receiver = $1 ORDER BY id",
        from,
    )
    if err != nil {
        log.Println("Query failed:", err)
        return
    }
    defer rows.Close()

    fmt.Println("\n--- Inbox ---")
    var count int
    for rows.Next() {
        var sender string
        var body string
        if err := rows.Scan(&sender, &body); err != nil {
            log.Println("Row scan failed:", err)
            continue
        }
        fmt.Printf("From %s: %s\n", sender, body)
        count++
    }
    if err := rows.Err(); err != nil {
        log.Println("Rows iteration error:", err)
    }
    if count == 0 {
        fmt.Println("No messages found.")
    }
    fmt.Println("------------\n")
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
			CheckInbox()
		}



	}

}

