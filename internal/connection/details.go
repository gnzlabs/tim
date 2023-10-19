package connection

import "crypto"

type Details struct {
	Address   string
	Port      int
	PublicKey crypto.PublicKey
}
