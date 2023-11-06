package secure_test

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/gnzlabs/tim/internal/connection/secure"
)

func getTestString(length int) (testString string) {
	testBuf := make([]byte, length)
	rand.Read(testBuf)
	testString = base64.StdEncoding.EncodeToString(testBuf)[:length]
	return
}

func getEstablishedConnections() (listener secure.Connection, client secure.Connection, err error) {
	if listener, err = secure.NewConnection(rand.Reader); err == nil {
		if client, err = secure.EstablishedConnection(rand.Reader, listener.PublicKey()); err == nil {
			err = listener.SetPeerKey(client.PublicKey())
		}
	}
	return
}

func TestHandshake(t *testing.T) {
	t.Log("Testing handshake....")
	if listener, client, err := getEstablishedConnections(); err != nil {
		t.Errorf("handshake failed: %s", err)
	} else {
		t.Logf("listener key: %s", base64.StdEncoding.EncodeToString(listener.PublicKey()[:]))
		t.Logf("client   key: %s", base64.StdEncoding.EncodeToString(client.PublicKey()[:]))
		t.Logf("listener shared key: %s", base64.StdEncoding.EncodeToString(listener.SharedKey()[:]))
		t.Logf("client   shared key: %s", base64.StdEncoding.EncodeToString(client.SharedKey()[:]))
		if *client.SharedKey() != *listener.SharedKey() {
			t.Error("shared key mismatch")
		}
	}
}

func TestClientEncryptListenerDecrypt(t *testing.T) {
	testString := getTestString(512)
	t.Logf("using test string with length %d bytes", len(testString))
	if listener, client, err := getEstablishedConnections(); err != nil {
		t.Errorf("handshake failed: %s", err)
	} else if ciphertext, err := client.Encrypt(testString); err != nil {
		t.Errorf("client encrypt failed: %s", err)
	} else if plaintext, err := listener.Decrypt(ciphertext); err != nil {
		t.Errorf("decrypt failed: %s", err)
	} else {
		t.Logf("ciphertext: %s", ciphertext)
		t.Logf("plaintext: %s", plaintext)
	}
}
