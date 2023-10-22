package secure

import "io"

type nonceHandler struct {
	rng        io.Reader
	usedNonces map[[24]byte]bool
}

func (n *nonceHandler) Contains(nonce *[24]byte) (exists bool) {
	_, exists = n.usedNonces[*nonce]
	return
}

func (n *nonceHandler) GenerateNew() *[24]byte {
	var nonce [24]byte
	n.rng.Read(nonce[:])
	for n.Contains(&nonce) {
		var nonce [24]byte
		n.rng.Read(nonce[:])
	}
	n.usedNonces[nonce] = true
	return &nonce
}
