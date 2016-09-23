package tar

import (
	"os"
	"testing"
)

func TestTarExtract(t *testing.T) {
	datamap, err := FileExtractor("i3lock-2.8-1-x86_64.pkg.tar")
	if err != nil {
		t.Fatal(err)
	}

	i := 0
	for k, v := range datamap {
		t.Logf("File path/name is: %s size is: %d\n", k, len(v.Data))
		i++
	}

	if i != 16 {
		t.Fatalf("Entry count should have been 16 instead is: %d\n", i)
	}
}

func TestTarNoExist(t *testing.T) {
	_, err := FileExtractor("nonesiste.tar")
	if err == nil {
		t.Fatal("Testing error condition gone!")
	}
}

func TestIsTarFileNoExist(t *testing.T) {
	ok, err := IsTarFile("nontarfile.tar")
	if ok {
		t.Fatal(err)
	}
	if err == nil {
		t.Fatal(err)
	}
}
func TestIsTarFile(t *testing.T) {
	ok, err := IsTarFile("i3lock-2.8-1-x86_64.pkg.tar")
	if !ok {
		t.Fatal(err)
	}
}
func TestTarBytesExtractor(t *testing.T) {
	f, err := os.Open("i3lock-2.8-1-x86_64.pkg.tar")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	buf := make([]byte, info.Size())
	n, err := f.Read(buf)
	if err != nil && int64(n) != info.Size() {
		t.Error("ops something went wrong")
	}

	_, err = Extractor(buf)
	if err != nil {
		t.Error(err)
	}

}
