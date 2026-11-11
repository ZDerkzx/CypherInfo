package srz

import (
	"bytes"
	"encoding/gob"
)

func Serialize(data any) ([]byte, error) {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Deserialize[T any](data []byte) (T, error) {
	var result T
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
