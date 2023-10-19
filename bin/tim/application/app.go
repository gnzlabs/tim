package application

import (
	"errors"
	"os"

	"github.com/gnzlabs/tim/internal/configuration"
	"github.com/gnzlabs/tim/internal/connection/client"
)

type Application struct {
	Configuration    *configuration.Configuration
	CurrentDirectory string
	Client           client.Client
}

func New(config *configuration.Configuration) (app *Application, err error) {
	if currentDirectory, err := os.Getwd(); err == nil {
		if client, err := client.New(&config.Host); err == nil {
			app = &Application{
				Configuration:    config,
				CurrentDirectory: currentDirectory,
				Client:           client,
			}
		}
	}
	return
}

func (app *Application) Run() (err error) {
	err = errors.New("not yet implemented")
	return
}
