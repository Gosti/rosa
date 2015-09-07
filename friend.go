package main

import (
	"crypto/rsa"
	"fmt"
	"strings"
)

//Friend contain all information to encrypt message for a friend, yes, without kidding
type Friend struct {
	Name      string
	PublicKey *rsa.PublicKey
}

//Local FriendList (if you have some friend)
var FriendList map[string]*Friend

func LoadFriends(filename string) error {
	filecontent, err := loadFile(filename)
	if err != nil {
		return err
	}
	friendList := strings.Split(string(filecontent), "\n")
	for _, friend := range friendList {
		s := strings.Split(friend, " ")
		if len(s) != 2 {
			break
		}
		FriendList[GetMD5Hash(s[1])] = &Friend{s[0], UnStringifyPublicKey(s[1])}
	}
	return nil
}

//Add friend f to FriendList In case I change simple list to a more complex type of data Binary tree or linked list will see later
func (f *Friend) Add() error {
	FriendList[GetMD5Hash(StringifyPublicKey(f.PublicKey))] = f
	return nil
}

func (f *Friend) Registrer(filename string) error {
	var content string = fmt.Sprintf("%s %v\n", f.Name, StringifyPublicKey(f.PublicKey))

	f.Add()

	err := appendFile(filename, []byte(content))

	return err
}

func (f *Friend) Remove(filename string) error {
<<<<<<< HEAD
	// Should I rewrote a shitty rewriter ? or do something a bit better ? like a grep... Seriously dunno
=======
	delete(FriendList, GetMD5Hash(StringifyPublicKey(f.PublicKey)))
>>>>>>> 1b66737c986c7de1e0b5fa037a7e2519d9361d05
	return nil
}

func (f *Friend) Encrypt(content []byte) ([]byte, error) {
	return Encrypt(content, f.PublicKey)
}

<<<<<<< HEAD
func SeekByName(name string) *Friend {
	for _, f := range FriendList {
		if f.Name == name {
			return f
		}
	}
	return nil
=======
func init() {
	FriendList = make(map[string]*Friend)
>>>>>>> 1b66737c986c7de1e0b5fa037a7e2519d9361d05
}
