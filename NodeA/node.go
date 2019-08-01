package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/libp2p/go-libp2p"
	"github.com/multiformats/go-multiaddr"
)

func addr2info(addrStr string) (*peer.AddrInfo, error) {

	addr, err := multiaddr.NewMultiaddr(addrStr)
	if err != nil {
		panic(err)
	}

	return peer.AddrInfoFromP2pAddr(addr)
}

func main() {

	var relayHost string

	flag.StringVar(&relayHost, "relay", "", "relay addr")

	flag.Parse()

	relayAddrInfo, err := addr2info(relayHost)

	// Zero out the listen addresses for the host, so it can only communicate
	// via p2p-circuit for our example
	h3, err := libp2p.New(context.Background(), libp2p.ListenAddrs(), libp2p.EnableRelay())
	if err != nil {
		panic(err)
	}

	if err := h3.Connect(context.Background(), *relayAddrInfo); err != nil {
		panic(err)
	}

	// Now, to test things, let's set up a protocol handler on h3
	h3.SetStreamHandler("/cats", func(s network.Stream) {
		fmt.Println("Meow! It worked!")
		s.Close()
	})

	fmt.Println("Node A ID: ", h3.ID())

	select {}
}
