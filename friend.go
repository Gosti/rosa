package main

import (
	"crypto/rsa"
	"fmt"
	"strings"
)

//Friend contain all information to encrypt and decrypt message
type Friend struct {
	Name      string
	PublicKey *rsa.PublicKey
}

//Local FriendList (if you have some friend)
var FriendList []*Friend

func LoadFriends(filename string) ([]Friend, error) {
	filecontent, err := loadFile(filename)
	if err != nil {
		return nil, err
	}
	friendList := strings.Split(string(filecontent), "\n")
	for _, friend := range friendList {
		fmt.Println(friend)
	}
	return nil, nil
}

//Add friend f to FriendList In case I change simple list to a more complex type of data Binary tree or linked list will see later
func (f *Friend) Add() error {
	FriendList = append(FriendList, f)
	return nil
}

func (f *Friend) Registrer(filename string) error {
	f.Add()

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

func (f *Friend) Remove(filename string) error {
	return nil
}

func (f *Friend) Encrypt(content []byte) ([]byte, error) {
	return Encrypt(content, f.PublicKey)
}
