# RoSA

Rosa is a fast and easy way to implement RSA in your go project.

### Version
0.0.1


### Feature
 - Generate Private and Public key and save them into a file
 - Manage Friend's Public key with a SSH like system (authorized_keys)

### Installation

```shell
go get github.com/mrgosti/rosa
```

### Usage

#### Basics

Import package
```go
import "github.com/mrgosti/rosa"
```

Generate a new pair of key
```go

```

Encrypt a message
```go
var key *rsa.PublicKey
//key is some *rsa.PublicKey you've load
message := []byte("Hello world!")
cryptedMessage, err := rosa.Encrypt(message, key)
```

#### Friends

Friends Struct
```go
type Friend struct {
	Name      string
	PublicKey *rsa.PublicKey
}
```

Load your friends (the already known key)
```go
rosa.LoadFriends(filename string)
// Will load Friends in rosa.FriendList
```

You can seek them by name
```go
f := rosa.SeekByName(name string) *Friend
```

Or by the md5 hash of their PublicKey
```go
rosa.FriendList[rosa.GetMD5Hash(*your key*)]
```

You can encrypt message for a friends
```go
f.Encrypt(content []byte) ([]byte, error)
// Just a shorthand for rosa.Encrypt
```