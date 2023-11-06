package secure

import "io"

type NonceHandler struct {
	Rng        io.Reader
	UsedNonces map[[24]byte]bool
}

func (n *NonceHandler) Contains(nonce *[24]byte) (exists bool) {
	_, exists = n.UsedNonces[*nonce]
	return
}

func (n *NonceHandler) Add(nonce [24]byte) {
	n.UsedNonces[nonce] = true
}

func (n *NonceHandler) GenerateNew() *[24]byte {
	var nonce [24]byte
	n.Rng.Read(nonce[:])
	for n.Contains(&nonce) {
		var nonce [24]byte
		n.Rng.Read(nonce[:])
	}
	n.Add(nonce)
	return &nonce
}
