# RoSA

RoSA is a fast and easy way to implement RSA in your go project.

### Version
0.0.1


### Feature
 - Generate Private and Public key and save them into a file
 - Manage Friend's Public key with a SSH like system (authorized_keys)

### Installation

```shell
go get github.com/mrgosti/rosa
```

### Tutorial / Example

```go
package main

import (
	"fmt"
	"github.com/mrgosti/rosa"
)

func main() {
	privateKey, publicKey, err := rosa.Generate("Example", false) // you generate a Key pair that you will use later (no need to save them)
	if err != nil {
		panic(err)
	}

	friend := &rosa.Friend{"Example", publicKey} // You create a new friends using your publickey as test

	cryptedMessage, err := friend.Encrypt([]byte("Hello World !")) // Same as doing rosa.Encrypt([]byte("Hello World !"), friend.PublicKey)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(cryptedMessage)

	msg, err := rosa.Decrypt(cryptedMessage, privateKey) // You decrypt it as well
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(msg))
}
```
