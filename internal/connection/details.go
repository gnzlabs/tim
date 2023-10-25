package connection

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Details struct {
	Address   net.IP
	Port      int
	PublicKey crypto.PublicKey
}

func (d *Details) PublicBytes() (keyBytes *[32]byte, err error) {
	if key, valid := d.PublicKey.(*[32]byte); !valid {
		err = errors.New("invalid key; type assertion failed")
	} else {
		keyBytes = key
	}
	return
}

func ParseDetails(ipAddress, portNumber, publicKey string) (details Details, err error) {
	if address := net.ParseIP(ipAddress); address == nil {
		err = fmt.Errorf("invalid IP address: %s", ipAddress)
	} else if port, e := strconv.Atoi(portNumber); e != nil {
		err = e
	} else if keyBytes, e := base64.StdEncoding.DecodeString(publicKey); e != nil {
		err = e
	} else if key, e := x509.ParsePKIXPublicKey(keyBytes); e != nil {
		err = e
	} else {
		details = Details{
			Address:   address,
			Port:      port,
			PublicKey: key,
		}
	}
	return
}
