package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// set path for dummy db
var path = "./Allmessage.json"

// configure the upgrader = an object with methods for taking a normal HTTP connection and upgrading it to a WebSocket
var upgrader = websocket.Upgrader{}

//
type Mail struct {
	ID          int    `json:"id"`
	DateCreated string `json:"created_date"`
	Message     string `json:"message"`
}

// for websocket
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Mail)              // broadcast channel

// Define HTTP request routes
func main() {

	// pull api (REST)
	r := mux.NewRouter()
	r.HandleFunc("/message", allMessage).Methods("GET")
	r.HandleFunc("/message", sendMessage).Methods("POST")

	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/", fs)

	// Configure websocket route PUSH API
	r.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	fmt.Println("starting web server at http://localhost:3100/")
	if err := http.ListenAndServe(":3100", r); err != nil {
		log.Fatal(err)
	}
}

func allMessage(w http.ResponseWriter, r *http.Request) {
	raw := getFIle(path)
	w.Header().Set("Content-Type", "application/json")

	w.Write(raw)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var mail Mail

	if err := json.NewDecoder(r.Body).Decode(&mail); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if mail.Message == "" {
		respondWithError(w, http.StatusBadRequest, "Error missing parameter message")
		return
	}

	now := time.Now()
	secs := now.Unix()

	mail.DateCreated = time.Unix(secs, 0).String()

	hasil := writeFile(mail)

	if hasil == false {
		respondWithError(w, http.StatusBadRequest, "Error While inserting message")
	} else {
		respondWithJson(w, http.StatusOK, map[string]string{"Success": "Successfuly insert message"})
	}

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func writeFile(message Mail) bool {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
		return false
	}

	var data_insert []Mail
	json.Unmarshal(raw, &data_insert)

	if len(data_insert) > 0 {

		last_slice := data_insert[len(data_insert)-1 : len(data_insert)][0]
		message.ID = last_slice.ID + 1

	} else {
		message.ID = 1
	}

	data_insert = append(data_insert, message)

	DataJson, _ := json.Marshal(data_insert)
	err = ioutil.WriteFile(path, DataJson, 0644)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return false
	}

	return true
}

func getFIle(file string) []byte {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(err.Error())
		os.Exit(1)

	}
	return raw
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Mail
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			// fmt.Println("socket clients ", ws, "close")
			break
		}
		// Send the newly received message to the broadcast channel

		now := time.Now()
		secs := now.Unix()

		msg.DateCreated = time.Unix(secs, 0).String()

		hasil := writeFile(msg)

		fmt.Println("messagenya", msg)

		if hasil {
			broadcast <- msg
		}
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				// fmt.Println("masuk error nh")
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
