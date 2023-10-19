package client

import (
	"crypto"
	"errors"

	"github.com/gnzlabs/tim/internal/command"
	"github.com/gnzlabs/tim/internal/connection"
)

type client struct {
	host       *connection.Details
	privateKey crypto.PrivateKey
	handlers   map[string]command.Handler
}

func New(host *connection.Details) (Client, error) {
	return nil, errors.New("not yet implemented")
}
