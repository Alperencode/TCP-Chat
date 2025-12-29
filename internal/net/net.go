package net

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func GetLocalIPAddress() (net.IP, error) {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		return nil, err
	}

	for _, r := range routes {
		if r.Dst == nil {
			continue
		}

		// first ip link with lowest metric
		link, err := netlink.LinkByIndex(r.LinkIndex)
		if err != nil {
			continue
		}

		// get the address list of the lowest metric link
		addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if addr.IP != nil && addr.IP.To4() != nil {
				return addr.IP, nil
			}
		}
	}

	return nil, fmt.Errorf("no local IPv4 address found")
}
