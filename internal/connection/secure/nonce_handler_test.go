package secure_test

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/gnzlabs/tim/internal/connection/secure"
)

func getNonceHandler() (nonceHandler *secure.NonceHandler) {
	return &secure.NonceHandler{
		Rng:        rand.Reader,
		UsedNonces: make(map[[24]byte]bool),
	}
}

func TestNonceHandler(t *testing.T) {
	nonceHandler := getNonceHandler()
	for i := 0; i < 4096; i++ {
		nonce := nonceHandler.GenerateNew()
		safeNonce := base64.StdEncoding.EncodeToString(nonce[:])
		t.Logf("generated nonce: %s", safeNonce)
		t.Logf("nonce exists: %t", nonceHandler.Contains(nonce))
	}
}
