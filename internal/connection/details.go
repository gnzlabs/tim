package connection

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"
)

type Details struct {
	Address   net.IP
	Port      int
	PublicKey crypto.PublicKey
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
