package pkgtoinfo

const (
	FILENAME     string = "%FILENAME%"
	NAME         string = "%NAME%"
	BASE         string = "%BASE%"
	VERSION      string = "%VERSION%"
	DESC         string = "%DESC%"
	GROUPS       string = "%GROUPS%"
	CSIZE        string = "%CSIZE%"
	ISIZE        string = "%ISIZE%"
	MD5SUM       string = "%MD5SUM%"
	SHA256SUM    string = "%SHA256SUM%"
	PGPSIG       string = "%PGPSIG%"
	URL          string = "%URL%"
	LICENSE      string = "%LICENSE%"
	ARCH         string = "%ARCH%"
	BUILDDATE    string = "%BUILDDATE%"
	PACKAGER     string = "%PACKAGER%"
	DEPENDS      string = "%DEPENDS%"
	OPTDEPENDS   string = "%OPTDEPENDS%"
	MAKEDEPENDS  string = "%MAKEDEPENDS%"
	CHECKDEPENDS string = "%CHECKDEPENDS%"
	PROVIDES     string = "%PROVIDES%"
)

type PackageInfo struct {
	Filename     string   `json:"filename"`
	Name         string   `json:"name"`
	Base         string   `json:"base"`
	Version      string   `json:"version"`
	Desc         string   `json:"desc"`
	Groups       string   `json:"groups"`
	Csize        string   `json:"csize"`
	Isize        string   `json:"isize"`
	Md5sum       string   `json:"md5sum"`
	Sha256sum    string   `json:"sha265sum"`
	Pgpsig       string   `json:"pgpsig"`
	Url          string   `json:"url"`
	License      string   `json:"license"`
	Arch         string   `json:"arch"`
	Builddate    string   `json:"builddate"`
	Packager     string   `json:"packager"`
	Depends      []string `json:"depends"`
	Makedepends  []string `json:"makedepends"`
	Checkdepends []string `json:"checkdepends"`
	Optdepends   []string `json:"optdepends"`
	Provides     []string `json:"provides"`
	//weight is neccessary when doing pkgdependency resolution
	Weight string `json:"weight"`
}
