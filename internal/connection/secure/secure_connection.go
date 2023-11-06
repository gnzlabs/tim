package secure

import (
	"errors"
	"io"

	"golang.org/x/crypto/nacl/box"
)

type secureConnection struct {
	privateKey    *[32]byte
	publicKey     *[32]byte
	peerPublicKey *[32]byte
	sharedKey     *[32]byte
	nonceHandler  *NonceHandler
}

func (c *secureConnection) establishConnection() (err error) {
	if c.peerPublicKey == nil {
		err = errors.New("can't establish connection; peer key not set")
	} else {
		var sharedKey [32]byte
		box.Precompute(&sharedKey, c.peerPublicKey, c.privateKey)
		c.sharedKey = &sharedKey
	}
	return
}

func (c *secureConnection) generateKey(rand io.Reader) (err error) {
	if c.privateKey != nil {
		err = errors.New("can't generate key; key already exists")
	} else {
		c.privateKey, c.publicKey, err = box.GenerateKey(rand)
	}
	return
}

func NewConnection(rand io.Reader) (connection Connection, err error) {
	c := &secureConnection{
		nonceHandler: &NonceHandler{
			Rng:        rand,
			UsedNonces: make(map[[24]byte]bool),
		},
	}
	err = c.generateKey(rand)
	return c, err
}

func EstablishedConnection(rand io.Reader, peerKey *[32]byte) (connection Connection, err error) {
	if connection, err = NewConnection(rand); err == nil {
		err = connection.SetPeerKey(peerKey)
	}
	return
}
