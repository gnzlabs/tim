package main

import (
	"fmt"
	"os"

	"github.com/gnzlabs/tim/internal/configuration"
)

func main() {
	var err error = nil
	if len(os.Args) != 4 {
		err = fmt.Errorf("expected arguments %s [host_address] [port] [public_key]", os.Args[0])
	}
	if config, e := configuration.New(os.Args[0], os.Args[1], os.Args[2], os.Args[3]); err == nil {
		application := Application{
			Configuration: config,
		}
		err = application.Run()
	} else {
		err = e
	}
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}
