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

func Generate(identifier string, save bool) (*rsa.PrivateKey, *rsa.PublicKey, error) {
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

	if save == true {
		savePrivateKey(privatekey, usr.HomeDir+"/.rosa/key2.priv")
		savePublicKey(publickey, identifier, usr.HomeDir+"/.rosa/key2.pub")
	}
	return privatekey, publickey, nil
}

func main() {
	usr, _ := user.Current()
	//Generate(usr.Username, true)
	_, err := LoadPrivateKey(usr.HomeDir + "/.rosa/key.priv")
	if err != nil {
		fmt.Println(err)
	}

	LoadFriends(usr.HomeDir + "/.rosa/friend_list")
<<<<<<< HEAD
	fmt.Printf("%+v\n", len(FriendList))
	me2 := SeekByName("gostimacbook")
=======
	fmt.Println(len(FriendList))
	FriendList["8b49905dcce57a634e18e386aa7f6b59"].Remove("fe")
	fmt.Println(len(FriendList))
	fmt.Println(FriendList["8b49905dcce57a634e18e386aa7f6b59"])
>>>>>>> 1b66737c986c7de1e0b5fa037a7e2519d9361d05
	// for i := 0; i < 40; i++ {
	// 	name := fmt.Sprintf("Test%d", i)
	// 	_, pub, _ := Generate(name, false)
	// 	f := &Friend{name, pub}
	// 	f.Registrer(usr.HomeDir + "/.rosa/friend_list")
	// }
	msg, _ := me2.Encrypt([]byte("Hello World!"))
	fmt.Println(string(msg))
	priv, err := LoadPrivateKey(usr.HomeDir + "/.rosa/key2.priv")
	decrypted, _ := Decrypt(msg, priv)

	fmt.Printf("%v\n", string(decrypted))
}
