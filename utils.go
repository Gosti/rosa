package rosa

import (
	"bytes"
	"encoding/gob"
)

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getInterface(bts []byte, data interface{}) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func save_file(title string, content []byte) error {
	title = string(bytes.Trim([]byte(title), "\x00"))
	content = bytes.Trim(content, "\x00")
	file, err := os.Create(title)
	if err != nil {
		return err
	}

	_, err = io.WriteString(file, string(content))
	if err != nil {
		return err
	}
	return nil
}
