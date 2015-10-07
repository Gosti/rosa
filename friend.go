package rosa

import (
	"crypto/rsa"
	"fmt"
	"regexp"
	"strings"
)

//Friend contain all information to encrypt message for a 'friend' or known host
type Friend struct {
	Name      string
	PublicKey *rsa.PublicKey
}

//FriendList is a map containing pointer on Friend and using the md5 hash (see rosa.GetMD5Hash()) of the PublicKey as an index
var FriendList map[string]*Friend

// LoadFriends help you to retrieve Friends saved in a file.
//
// The file have to be formatted as follow:
//
//	f.Name rosa.StringifyPublicKey(f.PublicKey)\n
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
		key, err := UnStringifyPublicKey(s[1])
		if err != nil {
			return nil
		}
		FriendList[GetMD5Hash(s[1])] = &Friend{s[0], key}
	}
	return nil
}

//Add friend f to FriendList In case I change simple list to a more complex type of data Binary tree or linked list will see later
func (f *Friend) Add() error {
	FriendList[GetMD5Hash(StringifyPublicKey(f.PublicKey))] = f
	return nil
}

// Register a Friend in the file that contains all the Friends, execute f.add too, Replace in case of duplicate
func (f *Friend) Register(filename string) error {
	var content = fmt.Sprintf("%s %v\n", f.Name, StringifyPublicKey(f.PublicKey))
	if already := SeekByName(f.Name); already != nil {
		err := already.Delete(filename)
		if err != nil {
			return err
		}
	}
	f.Add()
	err := appendFile(filename, []byte(content))

	return err
}

// Remove a Friend from the FriendList without deleting it from the file (see Delete for that behavior)
func (f *Friend) Remove() error {
	delete(FriendList, GetMD5Hash(StringifyPublicKey(f.PublicKey)))
	return nil
}

// Delete a Friend from both FriendList and Friend file
func (f *Friend) Delete(filename string) error {
	re := regexp.MustCompile("(" + f.Name + " " + regexp.QuoteMeta(StringifyPublicKey(f.PublicKey)) + "\n)")
	f.Remove()

	filecontent, err := loadFile(filename)
	if err != nil {
		return err
	}
	content := re.ReplaceAll(filecontent, []byte(""))
	saveFile(filename, content)
	return nil
}

// Encrypt is a short and for the rosa.Encrypt function using the Friend's public key directly
func (f *Friend) Encrypt(content []byte) ([]byte, error) {
	return Encrypt(content, f.PublicKey)
}

//SeekByName help you to find a friend by is name
func SeekByName(name string) *Friend {
	for _, f := range FriendList {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func init() {
	FriendList = make(map[string]*Friend)
}
