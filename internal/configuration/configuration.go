package configuration

import (
	"errors"

	"github.com/gnzlabs/tim/internal/connection"
)

type Configuration struct {
	BinaryName string
	Host       connection.Details
}

func New(binaryName, hostAddress string, port string, publicKey string) (config *Configuration, err error) {
	return nil, errors.New("not yet implemented")
}
