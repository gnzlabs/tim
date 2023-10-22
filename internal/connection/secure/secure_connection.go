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
	nonceHandler  *nonceHandler
}

func (c *secureConnection) establishConnection() (err error) {
	if c.peerPublicKey == nil {
		err = errors.New("can't establish connection; peer key not set")
	} else {
		box.Precompute(c.sharedKey, c.peerPublicKey, c.privateKey)
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

func NewConnection(rand io.Reader) (connection *secureConnection, err error) {
	connection = &secureConnection{
		nonceHandler: &nonceHandler{
			rng:        rand,
			usedNonces: make(map[[24]byte]bool),
		},
	}
	err = connection.generateKey(rand)
	return connection, err
}

func EstablishedConnection(rand io.Reader, peerKey *[32]byte) (connection *secureConnection, err error) {
	if connection, err = NewConnection(rand); err == nil {
		err = connection.SetPeerKey(peerKey)
	}
	return
}
