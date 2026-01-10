package client

import (
	"fmt"
	"log/slog"
	"net"
	localNet "tcp-chat/internal/net"
)

type Client struct {
	IP         net.IP
	Port       int
	Connection net.Conn
}

func NewClient(port int) *Client {
	ip, err := localNet.GetLocalIPAddress()
	if err != nil {
		return nil
	}

	return &Client{
		IP:   ip,
		Port: port,
	}
}

func (c *Client) Dial() error {
	host := c.IP.String() + ":" + fmt.Sprint(c.Port)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}

	c.Connection = conn
	slog.Info("TCP connection established", "Host", host)
	return nil
}

func (c *Client) Write(text string) error {
	_, err := c.Connection.Write([]byte(text))
	if err != nil {
		return err
	}

	slog.Info("Text successfully sent to server", "Text", text)
	return nil
}

func (c *Client) Close() error {
	return c.Connection.Close()
}
