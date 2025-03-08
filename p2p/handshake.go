package p2p

// Handshake func....
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
