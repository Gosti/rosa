package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	// "io/ioutil"
	"bytes"
	"encoding/gob"
	"io"
	"net"
	"os"
	"os/user"
)

type Friend struct {
	host      net.IP
	publicKey *rsa.PublicKey
	name      string
}

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getInterface(bts []byte, data interface{}) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func save_file(title string, content []byte) error {
	title = string(bytes.Trim([]byte(title), "\x00"))
	content = bytes.Trim(content, "\x00")
	file, err := os.Create(title)
	if err != nil {
		return err
	}

	_, err = io.WriteString(file, string(content))
	if err != nil {
		return err
	}
	return nil
}

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
	privatekey, _ := getBytes(privatekey)
	save_file(usr.HomeDir+"/.rosa/key.priv", toto)

	return privatekey, publickey, nil
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

func main() {
	usr, _ := user.Current()

	var wierd *rsa.PrivateKey

	privatekey, publickey, _ := Generate()
	toto
	getInterface(toto, &wierd)

	msg, _ := Encrypt([]byte("Hello world"), publickey)
	decrypted, _ := Decrypt(msg, wierd)

	fmt.Printf("%v\n", string(decrypted))
}
