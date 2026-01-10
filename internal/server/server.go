package server

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net"
	localNet "tcp-chat/internal/net"
)

type Server struct {
	IP       net.IP
	Port     int
	Listener net.Listener
}

func NewServer(port int) *Server {
	ip, err := localNet.GetLocalIPAddress()
	if err != nil {
		return nil
	}

	return &Server{
		IP:   ip,
		Port: port,
	}
}

func (s *Server) Listen() error {
	host := s.IP.String() + ":" + fmt.Sprint(s.Port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	s.Listener = listener
	slog.Info("Started listening incoming TCP connections", "Host", host)
	return err
}

func (s *Server) Accept() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			slog.Error("Error while accepting TCP connection", "Error", err.Error())
		}

		go func(c net.Conn) {
			var buf bytes.Buffer
			io.Copy(&buf, c)
			slog.Info("Received text", "Text", buf.String())
		}(conn)
	}
}

func (s *Server) Close() error {
	return s.Listener.Close()
}
