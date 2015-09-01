package main

import (
	"crypto/rsa"
	"fmt"
	"net"
	"strings"
)

type Friend struct {
	Host      net.IP
	PublicKey *rsa.PublicKey
	Name      string
}

var FriendList []*Friend

func LoadFriends(filename string) ([]Friend, error) {
	filecontent, err := load_file(filename)
	if err != nil {
		return nil, err
	}
	friendList := strings.Split(string(filecontent), "\n")
	for _, friend := range friendList {
		fmt.Println(friend)
	}
	return nil, nil
}

//In case I change simple list to a more complex type of data Binary tree or linked list will see later
func (f *Friend) Add() error {
	FriendList = append(FriendList, f)
	return nil
}

func (f *Friend) Registrer(filename string) error {
	f.Add()

	return nil
}

func (f *Friend) Remove(filename string) error {
	return nil
}

func (f *Friend) Encrypt(content []byte) ([]byte, error) {
	return Encrypt(content, f.PublicKey)
}
