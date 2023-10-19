package main

import (
	"errors"

	"github.com/gnzlabs/tim/internal/configuration"
)

type Application struct {
	Configuration    *configuration.Configuration
	CurrentDirectory string
}

func (app *Application) Run() (err error) {
	err = errors.New("not yet implemented")
	return
}
