package client

import "github.com/gnzlabs/tim/internal/command"

type Client interface {
	Connect() error
	Disconnect() error
	Send(string) error
	Register(command.Handler) error
}
