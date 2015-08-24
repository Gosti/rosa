package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"os/user"
)

func Decrypt(content []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
	md5hash := md5.New()
	label := []byte("")

	decryptedmsg, err := rsa.DecryptOAEP(md5hash, rand.Reader, privatekey, content, label)

	if err != nil {
		return nil, err
	}
	return decryptedmsg, nil
}

func Encrypt(content []byte, publickey *rsa.PublicKey) ([]byte, error) {
	md5hash := md5.New()
	label := []byte("")

	encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, publickey, content, label)
	if err != nil {
		return nil, err
	}
	return encryptedmsg, nil
}

func Generate() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	var publickey *rsa.PublicKey
	var privatekey *rsa.PrivateKey

	privatekey, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		return nil, nil, err
	}

	privatekey.Precompute()
	err = privatekey.Validate()

	if err != nil {
		return nil, nil, err
	}

	publickey = &privatekey.PublicKey

	return privatekey, publickey, nil
}

func isKeyAvailable() bool {

	usr, err := user.Current()
	if err != nil {
		return false
	}

	if _, err := os.Stat(usr.HomeDir + "/.rosa/key.priv"); err == nil {
		if _, err := os.Stat(usr.HomeDir + "/.rosa/key.pub"); err == nil {
			return true
		}
		return false
	}
	return false
}

func Retrieve() (*rsa.PrivateKey, error) {
	if isKeyAvailable() == false {
		privatekey, publickey, err := Generate()
		if err != nil {
			return nil, err
		}
		return
	} else {

	}
}

func main() {

	privatekey, publickey, _ := Generate()

	msg, _ := Encrypt([]byte("Hello world"), publickey)
	decrypted, _ := Decrypt(msg, privatekey)

	fmt.Printf("%v\n", decrypted)
	fmt.Println(isKeyAvailable())
}
