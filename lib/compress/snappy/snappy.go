package snappy

import "github.com/golang/snappy"

func Compress(data []byte) []byte {
	return snappy.Encode([]byte{}, data)
}

func Decompress(data []byte) ([]byte, error) {
	return snappy.Decode([]byte{}, data)
}
