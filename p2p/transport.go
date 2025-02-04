package p2p

// Peer is an interface that represents a remote node
type Peer interface {
}

// Transport is anyting that handles the communication
// between the nodes in the network.
// This can be of form (TCP, UDP, websockets)
type Transport interface {
	ListenAndAccept() error
}
