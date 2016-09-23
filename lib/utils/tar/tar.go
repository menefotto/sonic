package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"os"

	"github.com/sonic/lib/errors"
)

var ErrNotTarFile = errors.New("This is not a valid tar archive")

func IsTarFile(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, errors.Wrap(err)()
	}
	defer file.Close()

	ar := tar.NewReader(file)
	h, err := ar.Next()
	if h == nil {
		return false, ErrNotTarFile
	}
	return true, nil
}

type TarEntry struct {
	Data []byte
	Mode os.FileMode
}

func FileExtractor(filename string) (map[string]*TarEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	defer file.Close()

	return extractor(file)
}

func Extractor(tarbytes []byte) (map[string]*TarEntry, error) {
	buffer := bytes.NewReader(tarbytes)

	return extractor(buffer)
}

func extractor(reader io.Reader) (map[string]*TarEntry, error) {
	ar := tar.NewReader(reader)

	data := make(map[string]*TarEntry, 0)
	for {
		header, err := ar.Next()
		if err == io.EOF {
			break
		}
		if err != io.EOF && err != nil {
			return nil, errors.Wrap(err)()
		}

		content := make([]byte, header.Size)
		_, err = ar.Read(content)
		if err != io.EOF && err != nil {
			return nil, errors.Wrap(err)()
		}

		data[header.Name] = &TarEntry{Data: content, Mode: header.FileInfo().Mode()}
	}
	return data, nil

}
