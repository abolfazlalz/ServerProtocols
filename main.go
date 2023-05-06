package main

import (
	"ServerProtocols/net/http"
	"ServerProtocols/net/socket"
	"fmt"
)

func main() {
	httpServer := http.New()
	socketServer := socket.New()

	httpServer.StaticFolder("/", "public")
	socketServer.HandleHttp(httpServer)

	go func() {
		if err := httpServer.Listen(); err != nil {
			fmt.Printf("error during create http server: %v", err)
		}
	}()
	for {
		var cmd string
		_, err := fmt.Scanf("%s", &cmd)
		if err != nil {
			continue
		}

		if cmd == "send" {
			var id string
			var msg string
			fmt.Scanf("%s", &id)
			fmt.Scan(&msg)
			socketServer.Send(id, msg)
		}
	}
}
