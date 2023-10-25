package client

import (
	"crypto/rand"

	"github.com/gnzlabs/tim/internal/command"
	"github.com/gnzlabs/tim/internal/connection"
	"github.com/gnzlabs/tim/internal/connection/secure"
)

type SecureClient struct {
	host           *connection.Details
	messageHandler secure.Connection
	handlers       map[string]command.Handler
}

func New(host *connection.Details) (client Client, err error) {
	if publicKey, e := host.PublicBytes(); e != nil {
		err = e
	} else if messageHandler, e := secure.EstablishedConnection(rand.Reader, publicKey); e != nil {
		err = e
	} else {
		client = &SecureClient{
			host:           host,
			messageHandler: messageHandler,
			handlers:       make(map[string]command.Handler),
		}
	}
	return
}
