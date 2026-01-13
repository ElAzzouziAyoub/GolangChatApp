package main 

import(
	"fmt"
	"net/http"
	"sync"
	"encoding/json"
)


type Message struct{
	From string		`json:"body"`
	To string		`json:"body"`
	Body string `json:"body"`

}

var inboxes = make(map[string][]Message)
var mu sync.Mutex // protects inboxes


func handle(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Wassup"))
	fmt.Println("Connection established")
	switch r.Method {
	case http.MethodGet:
		
	case http.MethodPost:

		var msg Message
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// validation
		if msg.To == "" || msg.From == "" || msg.Body == "" {
			http.Error(w, "Missing fields", http.StatusBadRequest)
			return
		}

		// store message
		mu.Lock()
		inboxes[msg.To] = append(inboxes[msg.To], msg)
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message received"))


	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main(){

	http.HandleFunc("/",handle)
	http.ListenAndServe(":8080",nil)

}
