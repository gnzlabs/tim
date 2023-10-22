package configuration

import (
	"github.com/gnzlabs/tim/internal/connection"
)

type Configuration struct {
	BinaryName string
	Host       connection.Details
}

func New(binaryName, hostAddress string, port string, publicKey string) (config *Configuration, err error) {
	if connectionDetails, e := connection.ParseDetails(hostAddress, port, publicKey); e != nil {
		err = e
	} else {
		config = &Configuration{
			BinaryName: binaryName,
			Host:       connectionDetails,
		}
	}
	return
}
