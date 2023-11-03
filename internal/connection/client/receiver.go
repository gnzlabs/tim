package client

import (
	"bufio"
	"encoding/json"
	"net/textproto"

	"github.com/gnzlabs/tim/internal/command"
)

func (c *SecureClient) unpackMessage(packedMessage string) (message command.Message, err error) {
	if jsonMessage, e := c.messageHandler.Decrypt(packedMessage); e != nil {
		err = e
	} else {
		err = json.Unmarshal(jsonMessage, &message)
	}
	return
}

func (c *SecureClient) messageReceiver() {
	reader := bufio.NewReader(c.connection)
	textReader := textproto.NewReader(reader)
	for c.active {
		if packedMessage, err := textReader.ReadLine(); err != nil {
			c.log.Error("Read error", "error", err.Error())
		} else if message, err := c.unpackMessage(packedMessage); err != nil {
			c.log.Error("Parsing error", "error", err.Error())
		} else {
			c.commandChannel <- message
		}
	}
}
