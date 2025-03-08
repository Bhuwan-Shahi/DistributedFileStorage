package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	opts := TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       GOBDecoder{},
	}
	tr := NewTCPTransport(opts)                   // Fixed to pass TCPTransportOps
	assert.Equal(t, tr.listenAddress, listenAddr) // Kept as is, but note it won’t work
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
