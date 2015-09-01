package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"os/user"
)

func savePrivateKey(key *rsa.PrivateKey, filename string) error {

	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	err := save_file(filename, pemdata)
	if err != nil {
		return err
	}
	return nil
}

func savePublicKey(key *rsa.PublicKey, identifier string, filename string) error {
	str := fmt.Sprintf("%v%s", key.E, base64.StdEncoding.EncodeToString(key.N.Bytes()))
	err := save_file(filename, []byte(identifier+" "+str+"\n"))
	if err != nil {
		return err
	}
	return nil
}

func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	filecontent, err := load_file(filename)
	if err != nil {
		return nil, err
	}
	key, _ := pem.Decode(filecontent)
	if key == nil {
		return nil, errors.New("The file is not valid")
	}
	privatekey, err := x509.ParsePKCS1PrivateKey(key.Bytes)

	return privatekey, err
}

func isPrivKeyAvailable() bool {
	usr, err := user.Current()
	if err != nil {
		return false
	}

	if _, err := os.Stat(usr.HomeDir + "/.rosa/key.priv"); err == nil {
		return true
	}
	return false
}

func isPubKeyAvailable() bool {

	usr, err := user.Current()
	if err != nil {
		return false
	}

	if _, err := os.Stat(usr.HomeDir + "/.rosa/key.pub"); err == nil {
		return true
	}
	return false
}
