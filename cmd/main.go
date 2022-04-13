package main

import (
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) //connected clients

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	defer delete(clients, ws)

	// Make sure we close the connection when the function returns

	for {

		mt, message, err := ws.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}
		go sendMessages(message)
	}

}

func sendMessages(message []byte) {

	// Send it out to every client that is currently connected
	for ws := range clients {

		ws.WriteMessage(websocket.TextMessage, message)

	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	http.HandleFunc("/ws", handleConnections)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(viper.GetString("port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
