// Package rosa wrap the RSA package in a simplified form
package rosa

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
)

var PublicKeyPath string
var PrivateKeyPath string
var FriendListPath string

//Decrypt decrypt the given message with the given private key.
func Decrypt(content []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
	md5hash := md5.New()
	label := []byte("")

	decryptedmsg, err := rsa.DecryptOAEP(md5hash, rand.Reader, privatekey, content, label)

	if err != nil {
		return nil, err
	}
	return decryptedmsg, nil
}

//Encrypt Encrypt the given message with the given public key.
func Encrypt(content []byte, publickey *rsa.PublicKey) ([]byte, error) {
	md5hash := md5.New()
	label := []byte("")

	encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, publickey, content, label)
	if err != nil {
		return nil, err
	}
	return encryptedmsg, nil
}

//Generate is use to create a pair of keys (Private and Public) you can specify if you want to save them in a file,
// the path is defined by PrivateKeyPath and PublicKeyPath global variable
func Generate(identifier string, save bool) (*rsa.PrivateKey, *rsa.PublicKey, error) {
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

	if save == true {
		savePrivateKey(privatekey, PrivateKeyPath)
		savePublicKey(publickey, identifier, PublicKeyPath)
	}
	return privatekey, publickey, nil
}
