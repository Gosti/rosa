package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	// "io/ioutil"
	"os"
	"os/user"
)

type Friend struct {
	host      net.IP
	publicKey *rsa.PublicKey
	name      string
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

// TODO FIND HOW WRITE RSA PRIVATE KEY (AND PUBLIC I ASSUME) TO A FILE, IT MIGHT BE "FUNNY"
// func RetrievePrivate() *rsa.PrivateKey {
// 	usr, err := user.Current()
// 	if isPrivKeyAvailable() == false {
// 		return nil
// 	}
// 	content, err := ioutil.ReadFile(usr.HomeDir + "/.rosa/key.priv")
// 	if err != nil {
// 		return nil
// 	}
// 	return content.(rsa.PrivateKey)
// }
// func RetrieveFriend() []Friend

func main() {

	privatekey, publickey, _ := Generate()

	fmt.Printf("%v\n\n\n\n\n", []byte(privatekey))

	msg, _ := Encrypt([]byte("Hello world"), publickey)
	decrypted, _ := Decrypt(msg, privatekey)

	fmt.Printf("%v\n", decrypted)
}
