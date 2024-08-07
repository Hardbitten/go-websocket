package main

import (
	Server "main/server"
)

// var broadcast = make(chan Message)

// type Message struct {
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ws.Close()

// 	clients[ws] = true

// 	for {
// 		var msg Message
// 		err := ws.ReadJSON(&msg)

// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			delete(clients, ws)
// 			break
// 		}

// 		broadcast <- msg
// 	}
// }

// func handleMessages() {
// 	for {
// 		msg := <-broadcast
// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

func main() {
	srv := Server.Server{}
	srv.StartServer()

}
