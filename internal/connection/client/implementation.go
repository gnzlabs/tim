package client

import (
	"errors"
	"fmt"
	"net"

	"github.com/gnzlabs/tim/internal/command"
)

// Connect implements Client.Connect for client
func (c *SecureClient) Connect() (err error) {
	if c.connection, err = net.Dial("tcp", c.host.ConnectionString()); err == nil {
		c.active = true
	}
	return
}

// Disconnect implements Client.Disconnect for client
func (c *SecureClient) Disconnect() error {
	return errors.New("not yet implemented")
}

// Send implements Client.Send for client
func (c *SecureClient) Send(message string) error {
	return errors.New("not yet implemented")
}

// Register implements Client.Register for client
func (c *SecureClient) Register(handler command.Handler) (err error) {
	if _, exists := c.handlers[handler.Name()]; exists {
		err = fmt.Errorf("handler %s already registered", handler.Name())
	} else {
		c.handlers[handler.Name()] = handler
	}
	return
}
