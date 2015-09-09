package rosa

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

//GetMD5Hash help you to get the key for the map FriendList
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func saveFile(filename string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, string(content))
	if err != nil {
		return err
	}
	return nil
}

// I know this function is pretty stupid but I really want All I/O manipulation and include in this file
func loadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func appendFile(filename string, content []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Write(content); err != nil {
		return err
	}
	return nil
}
