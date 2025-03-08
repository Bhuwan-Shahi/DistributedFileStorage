package p2p

// Meassage holds any arbiratary data that is being sent over the
// each transport between two nodes in the network.
type Messgae struct {
	Payload []byte
}
