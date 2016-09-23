package xz

import (
	"os"
	"testing"
)

func TestXzFileDecompressionByte(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}
func TestXzFileDecompressionWrong(t *testing.T) {
	f, err := os.Open("malformed.txt.xz")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}

func TestXzFileDecompressionNoEsist(t *testing.T) {
	f, err := os.Open("nofilehere.txt")
	if err == nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err == nil {
		t.Fatal(err)
		return
	} else {
		return
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Decompress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}
func TestXzFileDecompressionString(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(string(dat))
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}

func TestXzDeXzCompressBytes(t *testing.T) {
	f, err := os.Open("test.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestXzDeXzCompressBytesMalformed(t *testing.T) {
	f, err := os.Open("malformed.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err == nil {
		t.Log(err)
	}

	if len(b) == 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}
func TestXzDeXzCompressNoFile(t *testing.T) {
	f, err := os.Open("")
	if err == nil {
		t.Fatal(err)
	}

	info, err := f.Stat()
	if err == nil {
		t.Error(err)
		return
	} else {
		return
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestDeXzCompressWrongInvalidInput(t *testing.T) {
	data := ""
	_, err := Compress(data)
	if err == nil {
		t.Log(err)
	}

}
func TestXzDeXzCompressString(t *testing.T) {
	f, err := os.Open("test.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(string(data))
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestFileXzDeXzCompress(t *testing.T) {
	data, err := DecompressFile("test.txt.xz")
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1203 {
		t.Fatal("DeXzCompression went wrong, size is: ", len(data))
	}
}

func TestFileXzXzCompress(t *testing.T) {
	err := CompressFile("test.txt", "test.xz")
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat("test.xz")
	if err != nil {
		t.Fatal(err)
	}

	if info.Size() != int64(772) {
		t.Fatal("XzCompression went wrong, size missmatch: ", info.Size())
	}

	os.Remove("test.xz")
}

func TestFileXzXzCompressWrong(t *testing.T) {
	err := CompressFile("test.txt", "/var/cache/test.xz")
	if err == nil {
		t.Fatal(err)
	}

	_, err = os.Stat("/var/cache/test.xz")
	if err == nil {
		t.Fatal(err)
	}

}

func TestFileXzXzCompressWrongInput(t *testing.T) {
	err := CompressFile("boh", "test.xz")
	if err == nil {
		t.Fatal(err)
	}

	_, err = os.Stat("test.xz")
	if err == nil {
		t.Fatal(err)
	}

}
