package client

import (
	"errors"
	"fmt"

	"github.com/gnzlabs/tim/internal/command"
)

// Connect implements Client.Connect for client
func (c *client) Connect() error {
	return errors.New("not yet implemented")
}

// Disconnect implements Client.Disconnect for client
func (c *client) Disconnect() error {
	return errors.New("not yet implemented")
}

// Send implements Client.Send for client
func (c *client) Send(message string) error {
	return errors.New("not yet implemented")
}

// Register implements Client.Register for client
func (c *client) Register(handler command.Handler) (err error) {
	if _, exists := c.handlers[handler.Name()]; exists {
		err = fmt.Errorf("handler %s already registered", handler.Name())
	} else {
		c.handlers[handler.Name()] = handler
	}
	return
}
