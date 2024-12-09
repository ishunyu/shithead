package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error %s when reading message from client", err)
			return
		}
		if mt == websocket.BinaryMessage {
			err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
			return
		}
		log.Printf("Receive message %s", string(message))
		command := strings.Trim(string(message), "\n")
		if command == "start" {
			err = c.WriteMessage(websocket.TextMessage, []byte("Game starting..."))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
			return
		}
	}

}

func main() {
	var origins = []string{"null", "http://localhost:8080"}
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{ // Resolve cross-domain problems
			CheckOrigin: func(r *http.Request) bool {
				var origin = r.Header.Get("origin")
				for _, allowOrigin := range origins {
					if origin == allowOrigin {
						return true
					}
				}
				return false
			}},
	}
	http.Handle("/", webSocketHandler)
	log.Print("Starting shithead server...\n")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
