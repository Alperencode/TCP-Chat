package server

import (
	"net"
	"sync"
)

type Server struct {
	listener net.Listener
	clients  map[string]net.Conn
	mu       sync.Mutex
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	server := Server{
		listener: listener,
		clients:  make(map[string]net.Conn),
		mu:       sync.Mutex{},
	}

	return &server, nil
}
