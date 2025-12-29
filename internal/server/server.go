package server

import (
	"net"
	localNet "tcp-chat/internal/net"
)

type Server struct {
	IP net.IP
}

func NewServer() *Server {
	ip, err := localNet.GetLocalIPAddress()
	if err != nil {
		return nil
	}

	return &Server{
		IP: ip,
	}
}
