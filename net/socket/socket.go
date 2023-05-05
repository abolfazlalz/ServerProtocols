package socket

import (
	"ServerProtocols/net/http"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"strconv"
)

type Client struct {
	*websocket.Conn
	Id string
}

type Server struct {
	connections map[string]*Client
	count       uint
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func New() *Server {
	return &Server{
		connections: make(map[string]*Client),
		count:       0,
	}
}

func (s *Server) handleWs(ws *websocket.Conn) {
	fmt.Println("New incoming web socket message: ", ws.RemoteAddr())

	s.count += 1

	id := getMD5Hash(strconv.Itoa(int(s.count)))

	client := &Client{
		Conn: ws,
		Id:   id,
	}

	s.connections[id] = client

	s.readLoop(client)
}

func (s *Server) readLoop(client *Client) {
	buf := make([]byte, 1024)

	for {
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				delete(s.connections, client.Id)
				break
			}
			fmt.Println("Read error:", err)
			continue
		}

		msg := string(buf[:n])

		s.broadcast(fmt.Sprintf("%s - %s", client.Id, msg))
	}
}

func (s *Server) broadcast(msg string) {
	for _, client := range s.connections {
		go func(ws *Client) {
			if _, err := ws.Write([]byte(msg)); err != nil {
				fmt.Printf("Write error, %v", err)
			}
		}(client)
	}
}

func (s *Server) HandleHttp(http *http.Server) {
	http.Handle("/ws", websocket.Handler(s.handleWs))
}
