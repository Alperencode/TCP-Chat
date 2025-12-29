package main

import (
	"flag"
	"log/slog"
	"tcp-chat/internal/client"
	"tcp-chat/internal/server"
)

func main() {
	typeFlagPtr := flag.String("type", "client", "start client or server")
	flag.Parse()

	switch *typeFlagPtr {
	case "client":
		client := client.NewClient()
		slog.Info("Client created.", "Client", client)
	case "server":
		server := server.NewServer()
		slog.Info("Server created.", "Server", server)
	default:
		slog.Error("Type should be a client or server", "Type", *typeFlagPtr)
	}
}
