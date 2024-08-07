package server

import (
	"log"
	"main/event"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[int]*websocket.Conn)
var DataChannel = make(chan event.Data)
var ServerUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Server struct {
	LastUserId int
}

func (s *Server) StartServer() {

	// initiate
	s.LastUserId = 0

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", s.HandleConnections)

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
