package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"bytes"
)


type Message struct{
	From string		`json:"body"`
	To string		`json:"body"`
	Body string `json:"body"`

}



var from string
var inbox []string


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
				log.Fatal(err)
			}


			resp, err := http.Post(
				"http://localhost:8080/",
				"application/json",
				bytes.NewBuffer(data),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

}
func CheckInbox(){

}

func main(){
	fmt.Print("Enter your username :")
	fmt.Scanln(&from)
	fmt.Printf("\nWelcome %s To the GolangChatApp \n")


	for {
		var choice int 

		fmt.Println("Choose an Option :")
		fmt.Println("1-Send a message ")
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

