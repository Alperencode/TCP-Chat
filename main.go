package main

import (
	"flag"
	"fmt"
)

func main() {
	typeFlagPtr := flag.String("type", "client", "start client or server")
	flag.Parse()

	switch *typeFlagPtr {
	case "client":
	case "server":
	default:
		fmt.Println("Type should be a client or server")
	}
}
