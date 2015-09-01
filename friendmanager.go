package main

type Friend struct {
	host      net.IP
	publicKey *rsa.PublicKey
	name      string
}

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

func (f *Friend) add(filename string) error {

}
