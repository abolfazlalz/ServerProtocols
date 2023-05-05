package main

import (
	"ServerProtocols/net/http"
	"ServerProtocols/net/socket"
	"fmt"
)

func main() {
	httpServer := http.New()
	socketServer := socket.New()

	socketServer.HandleHttp(httpServer)

	if err := httpServer.Listen(); err != nil {
		fmt.Printf("error during create http server: %v", err)
	}
}
