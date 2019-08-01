package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
)

func main() {

	listenPort := flag.Int("l", 53100, "wait for incoming connections")
	flag.Parse()

	h2, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *listenPort)),
		libp2p.EnableRelay(circuit.OptHop))
	if err != nil {
		panic(err)
	}

	for _, ips := range h2.Addrs() {
		fmt.Printf("%s/p2p/%s\n", ips, h2.ID())
	}

	fmt.Printf("%s/p2p/%s\n", h2.Addrs()[len(h2.Addrs())-1], h2.ID())

	select {}

}
