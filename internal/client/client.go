package client

import (
	"net"
	localNet "tcp-chat/internal/net"
)

type Client struct {
	IP net.IP
}

func NewClient() *Client {
	ip, err := localNet.GetLocalIPAddress()
	if err != nil {
		return nil
	}

	return &Client{
		IP: ip,
	}
}
