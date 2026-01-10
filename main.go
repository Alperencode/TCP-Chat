package main

import (
	"flag"
	"log/slog"
	"tcp-chat/internal/client"
	"tcp-chat/internal/server"
)

func main() {
	typeFlagPtr := flag.String("type", "client", "start client or server")
	portFlagPtr := flag.Int("port", 8080, "port to connect")
	flag.Parse()

	switch *typeFlagPtr {
	case "client":
		client := client.NewClient(*portFlagPtr)
		slog.Info("Client created.", "Client", client)

		err := client.Dial()
		if err != nil {
			slog.Error("Error while establishing TCP connection", "Error", err.Error())
		}
		defer client.Close()

		client.Write("Hello World!")
	case "server":
		server := server.NewServer(*portFlagPtr)
		slog.Info("Server created.", "Server", server)

		err := server.Listen()
		if err != nil {
			slog.Error("Error while trying to listen TCP connections", "Error", err.Error())
		}
		defer server.Close()

		server.Accept()
	default:
		slog.Error("Type should be a client or server", "Type", *typeFlagPtr)
	}
}
