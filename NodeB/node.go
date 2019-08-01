package main

import (
	"context"
	"flag"
	"fmt"

	swarm "github.com/libp2p/go-libp2p-swarm"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/multiformats/go-multiaddr"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
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
	var dialID string

	flag.StringVar(&relayHost, "relay", "", "relay addr")
	flag.StringVar(&dialID, "dial", "", "dial Node ID")

	flag.Parse()

	relayAddrInfo, err := addr2info(relayHost)

	h1, err := libp2p.New(context.Background(), libp2p.EnableRelay(circuit.OptDiscovery))
	if err != nil {
		panic(err)
	}

	if err := h1.Connect(context.Background(), *relayAddrInfo); err != nil {
		panic(err)
	}

	dialNodeID, err := peer.IDB58Decode(dialID)

	if err != nil {
		panic(err)
	}

	_, err = h1.NewStream(context.Background(), dialNodeID, "/cats")
	if err == nil {
		fmt.Println("Didnt actually expect to get a stream here. What happened?")
		return
	}
	fmt.Println("Okay, no connection from h1 to h3: ", err)
	fmt.Println("Just as we suspected")

	h1.Network().(*swarm.Swarm).Backoff().Clear(dialNodeID)

	ma.SwapToP2pMultiaddrs()
	relayaddr, err := ma.NewMultiaddr(relayHost + "/p2p-circuit/p2p/" + dialNodeID.Pretty())

	if err != nil {
		panic(err)
	}

	h3relayInfo := peer.AddrInfo{
		ID:    dialNodeID,
		Addrs: []ma.Multiaddr{relayaddr},
	}

	if err := h1.Connect(context.Background(), h3relayInfo); err != nil {
		panic(err)
	}

	// Woohoo! we're connected!
	s, err := h1.NewStream(context.Background(), dialNodeID, "/cats")
	if err != nil {
		fmt.Println("huh, this should have worked: ", err)
		return
	}

	s.Read(make([]byte, 1)) // block until the handler closes the stream

	fmt.Println("end")

}
