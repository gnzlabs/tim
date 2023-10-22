package secure

import "errors"

// PublicKey implements SecureConnection.PublicKey for secureConnection
func (c *secureConnection) PublicKey() (publicKey [32]byte) {
	if c.publicKey != nil {
		publicKey = *c.publicKey
	}
	return
}

// SetPeerKey implements SecureConnection.SetPeerKey for secureConnection
func (c *secureConnection) SetPeerKey(peerKey *[32]byte) (err error) {
	if c.peerPublicKey != nil {
		err = errors.New("can't set peer key; key already exists")
	} else {
		c.peerPublicKey = peerKey
		err = c.establishConnection()
	}
	return
}

// Encrypt implements SecureConnection.Encrypt for secureConnection
func (c *secureConnection) Encrypt(plaintext string) (ciphertext string, err error) {
	return "", errors.New("not yet implemented")
}

// Decrypt implements SecureConnection.Decrypt for secureConnection
func (c *secureConnection) Decrypt(ciphertext string) (plaintext string, err error) {
	return "", errors.New("not yet implemented")
}
