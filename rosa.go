package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
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

func Generate(identifier string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	var publickey *rsa.PublicKey
	var privatekey *rsa.PrivateKey

	usr, err := user.Current()
	privatekey, err = rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		return nil, nil, err
	}

	privatekey.Precompute()
	err = privatekey.Validate()

	if err != nil {
		return nil, nil, err
	}

	publickey = &privatekey.PublicKey

	savePrivateKey(privatekey, usr.HomeDir+"/.rosa/key.priv")
	savePublicKey(publickey, identifier, usr.HomeDir+"/.rosa/key.pub")
	return privatekey, publickey, nil
}

func main() {
	usr, _ := user.Current()
	Generate("mrgosti")
	_, err := LoadPrivateKey(usr.HomeDir + "/.rosa/key.priv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", usr.Username)

	// msg, _ := Encrypt([]byte("Hello world"), publickey)
	// decrypted, _ := Decrypt(msg, wierd)

	// fmt.Printf("%v\n", string(decrypted))
}
