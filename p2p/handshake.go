package p2p

// Handshake func....
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
