# go-libp2p-relay
An example that demonstrates how to use circuit relay.


This is an example that quickly shows how to use the `github.com/libp2p/go-libp2p-circuit` package to  build and use relay peer.

This example is based on the [libp2p/go-libp2p-examples/relay](https://github.com/libp2p/go-libp2p-examples/tree/master/relay) modification, and the difference from the source code is that this example splits a file in the original text into 3  peers. And a particularly important change is


```go
// code in https://github.com/libp2p/go-libp2p-examples/blob/b7ac9e91865656b3ec13d18987a09779adad49dc/relay/main.go#L68
// Creates a relay address
relayaddr, err := ma.NewMultiaddr("/p2p-circuit/ipfs/" + h3.ID().Pretty())
if err != nil {
	panic(err)
}

relayaddr, err := ma.NewMultiaddr(fmt.Printf("%s/p2p/%s\n", h2.Addrs()[len(h2.Addrs())-1], h2.ID()) + "/p2p-circuit/ipfs/" + h3.ID().Pretty())
if err != nil {
	panic(err)
}



```


NodeA and NodeB are behind NAT, reverse proxies, firewalls and/or simply don't able to establish a direct connection to each other.

Now we can let them establish a connection by step:

```

NodeA -> NodeRelay

NodeB -> NdoeRelay

NodeA -> NodeB
```

NodeA and NodeB can be deployed separately on servers on defferent networks.

NodeRelay must be deployed on server on a publicly accessible ip:port.

## Usage

NodeA, NodeB, NodeRelay can be deployed separately on different network servers.

### Run Node Relay

```
> cd go-libp2p-relay/NodeRelay
> go run node.go

/ip4/127.0.0.1/tcp/53100/p2p/QmUtH7d63G1jhbfdqUzQ7HeCkxVG2GVukTjkcs5ZTDLx6N
/ip4/xxx.xxx.xxx.xxx/tcp/53100/p2p/QmUtH7d63G1jhbfdqUzQ7HeCkxVG2GVukTjkcs5ZTDLx6N

```

xxx.xxx.xxx.xxx is your server public IP


### Run Node A

```
> cd go-libp2p-relay/NodeA
> go run node.go -relay /ip4/xxx.xxx.xxx.xxx/tcp/53100/p2p/QmUtH7d63G1jhbfdqUzQ7HeCkxVG2GVukTjkcs5ZTDLx6N

Node A ID:  Qmb1AQPRCvFTz1rj4CqNuFs6NsEiJcpnkFRzzcZRAVfVgD

```

### Run Node B

```
> cd go-libp2p-relay/NodeB
> go run node.go -relay /ip4/xxx.xxx.xxx.xxx/tcp/53100/p2p/QmUtH7d63G1jhbfdqUzQ7HeCkxVG2GVukTjkcs5ZTDLx6N -dial Qmb1AQPRCvFTz1rj4CqNuFs6NsEiJcpnkFRzzcZRAVfVgD

Okay, no connection from h1 to h3:  failed to dial Qmb1AQPRCvFTz1rj4CqNuFs6NsEiJcpnkFRzzcZRAVfVgD: no addresses
Just as we suspected
end

```

-dial Qmb1AQPRCvFTz1rj4CqNuFs6NsEiJcpnkFRzzcZRAVfVgD  
-dial NodeA ID that can be found in the terminal output of server Node A

If all goes well, you will see the output `Meow! It worked!` at the terminal of NodeA.


## Thanks

[libp2p/go-libp2p-examples](https://github.com/libp2p/go-libp2p-examples/tree/master/relay)  
[@raulk](https://github.com/raulk)  
[@upperwal](https://github.com/upperwal)  



