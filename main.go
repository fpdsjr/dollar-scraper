package main

import (
	"fmt"
	"net/http"

	"github.com/fpdsjr/dolar-scraper/scraper"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSDollar(ws *websocket.Conn) {
	s.conns[ws] = true

	for {
		scraper.ScraperDollar()
		websocket.Message.Send(ws, scraper.DollarPrice)
	}
}

func main() {
	s := NewServer()

	fmt.Println("Websocket server running on port 8080")
	http.Handle("/ws", websocket.Handler(s.handleWSDollar))
	http.ListenAndServe(":8080", nil)
}
