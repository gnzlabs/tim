package client

import (
	"crypto/rand"
	"log/slog"
	"net"

	"github.com/gnzlabs/tim/internal/command"
	"github.com/gnzlabs/tim/internal/connection"
	"github.com/gnzlabs/tim/internal/connection/secure"
)

type SecureClient struct {
	active         bool
	connection     net.Conn
	commandChannel chan command.Message
	handlers       map[string]command.Handler
	host           *connection.Details
	log            *slog.Logger
	messageHandler secure.Connection
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
