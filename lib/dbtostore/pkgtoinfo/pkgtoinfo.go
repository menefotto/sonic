// this package offers to function FileToMap
// and StringToMap they are specialized functions
// that map the depends and desc file from each
// arch linux package to series of key and values
// and return a map[string][]string

package pkgtoinfo

import (
	"io/ioutil"
	"strings"

	"github.com/sonic/lib/errors"
)

func FileToPkgInfo(name string) (*PackageInfo, error) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	return toPackageInfo(string(content)), nil
}

func StringToPkgInfo(data string) *PackageInfo {
	return toPackageInfo(data)
}

func toPackageInfo(str string) *PackageInfo {
	info := &PackageInfo{}
	text := strings.Split(str, "\n")

	for i := 0; i < len(text); i++ {
		if len(text[i]) > 0 && strings.HasPrefix(text[i], "%") {
			singleSectionMapper(i, info, text)
		}
	}

	return info

}

// ugliest code I have ever written, for now that's good enough, soon I will write
// a better version.

func singleSectionMapper(i int, info *PackageInfo, text []string) {

	if idx := i + 1; idx < len(text) {
		switch {
		case text[i] == FILENAME:
			info.Filename = text[idx]
		case text[i] == NAME:
			info.Name = text[idx]
		case text[i] == BASE:
			info.Base = text[idx]
		case text[i] == VERSION:
			info.Version = text[idx]
		case text[i] == DESC:
			info.Desc = text[idx]
		case text[i] == GROUPS:
			info.Groups = text[idx]
		case text[i] == CSIZE:
			info.Csize = text[idx]
		case text[i] == ISIZE:
			info.Isize = text[idx]
		case text[i] == MD5SUM:
			info.Md5sum = text[idx]
		case text[i] == SHA256SUM:
			info.Sha256sum = text[idx]
		case text[i] == PGPSIG:
			info.Pgpsig = text[idx]
		case text[i] == URL:
			info.Url = text[idx]
		case text[i] == LICENSE:
			info.License = text[idx]
		case text[i] == ARCH:
			info.Arch = text[idx]
		case text[i] == BUILDDATE:
			info.Builddate = text[idx]
		case text[i] == PACKAGER:
			info.Packager = text[idx]
		case text[i] == MAKEDEPENDS:
			multiSectionMapper(i, info, text, "M")
		case text[i] == DEPENDS:
			multiSectionMapper(i, info, text, "D")
		case text[i] == CHECKDEPENDS:
			multiSectionMapper(i, info, text, "C")
		case text[i] == OPTDEPENDS:
			multiSectionMapper(i, info, text, "O")
		case text[i] == PROVIDES:
			multiSectionMapper(i, info, text, "P")

		}
	}
}

func multiSectionMapper(i int, info *PackageInfo, text []string, t string) {
	for idx := i + 1; idx < len(text) && !strings.HasPrefix(text[idx], "%"); idx++ {
		if len(text[idx]) > 0 {
			switch {
			case t == "M":
				info.Makedepends = append(info.Makedepends, text[idx])
			case t == "D":
				info.Depends = append(info.Depends, text[idx])
			case t == "C":
				info.Checkdepends = append(info.Checkdepends, text[idx])
			case t == "O":
				info.Optdepends = append(info.Optdepends, text[idx])
			case t == "P":
				info.Provides = append(info.Provides, text[idx])
			}
		}
	}
}
