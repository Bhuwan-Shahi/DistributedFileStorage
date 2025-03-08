package main

import (
	"log"

	"github.com/Bhuwan-Shahi/DistributedFilsStorage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOps{
		ListenAddr:    "3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

}
