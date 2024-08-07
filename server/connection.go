package server

import (
	"log"
	"main/event"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := ServerUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.LastUserId += 1
	currentUser := s.LastUserId
	Clients[currentUser] = ws
	defer s.HandleDisconnect(ws, currentUser)

	log.Println("user[", currentUser, "] connected")

	go s.HandleIncomingData(ws, currentUser)
	for {
		data := event.Data{UserId: currentUser}
		err := ws.ReadJSON(&data)
		if err != nil {
			log.Printf("error: %v", err)
			// delete(clients, ws)
			break
		}
		DataChannel <- data
	}

}

func (s *Server) HandleIncomingData(ws *websocket.Conn, currentUser int) {
	for {
		msg := <-DataChannel

		for client := range Clients {
			socket := Clients[client]
			err := socket.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
			}

		}
	}
}

func (s *Server) HandleDisconnect(ws *websocket.Conn, userId int) {
	delete(Clients, userId)
	log.Println("user[", userId, "] disconnected")
	ws.Close()
}
