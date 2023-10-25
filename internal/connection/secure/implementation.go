package secure

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/nacl/box"
)

// PublicKey implements Connection.PublicKey for secureConnection
func (c *secureConnection) PublicKey() (publicKey [32]byte) {
	if c.publicKey != nil {
		publicKey = *c.publicKey
	}
	return
}

// SetPeerKey implements Connection.SetPeerKey for secureConnection
func (c *secureConnection) SetPeerKey(peerKey *[32]byte) (err error) {
	if c.peerPublicKey != nil {
		err = errors.New("can't set peer key; key already exists")
	} else {
		c.peerPublicKey = peerKey
		err = c.establishConnection()
	}
	return
}

// Encrypt implements Connection.Encrypt for secureConnection
func (c *secureConnection) Encrypt(plaintext string) (ciphertext string, err error) {
	plaintextBytes := []byte(plaintext)
	nonce := c.nonceHandler.GenerateNew()
	ciphertextBytes := make([]byte, len(plaintextBytes)+box.Overhead)
	box.SealAfterPrecomputation(ciphertextBytes, plaintextBytes, nonce, c.sharedKey)
	packedBytes := append(nonce[:], ciphertextBytes...)
	ciphertext = base64.StdEncoding.EncodeToString(packedBytes)
	return
}

// Decrypt implements Connection.Decrypt for secureConnection
func (c *secureConnection) Decrypt(ciphertext string) (plaintext string, err error) {
	var nonce [24]byte
	if ciphertextBytes, e := base64.StdEncoding.DecodeString(ciphertext); e != nil {
		err = e
	} else if len(ciphertextBytes) < 24+box.Overhead {
		err = errors.New("decryption error: invalid length")
	} else {
		copy(nonce[:], ciphertextBytes[:24])
		if c.nonceHandler.Contains(&nonce) {
			err = errors.New("decryption error: nonce reuse")
		} else {
			plaintextBytes := make([]byte, len(ciphertextBytes[:24])-box.Overhead)
			if plaintextBytes, valid := box.OpenAfterPrecomputation(plaintextBytes, ciphertextBytes[:24], &nonce, c.sharedKey); !valid {
				err = errors.New("decryption error: validation failed")
			} else {
				c.nonceHandler.Add(nonce)
				plaintext = string(plaintextBytes)
			}
		}
	}
	return
}
