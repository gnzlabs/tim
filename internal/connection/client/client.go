package client

import (
	"github.com/gnzlabs/tim/internal/command"
	"github.com/gnzlabs/tim/internal/connection"
)

type client struct {
	host            connection.Details
	commandHandlers map[string]command.Handler
}
