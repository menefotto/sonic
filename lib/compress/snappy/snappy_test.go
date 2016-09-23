package snappy

import "testing"

func TestSnappyCompressDecompress(t *testing.T) {
	str := "ciao carlo"
	res := Compress([]byte(str))
	dec, err := Decompress(res)
	if err != nil {
		t.Fatal(err)
	}
	if str != string(dec) {
		t.Fatal("something happened in str")
	}

}
