package crypto

import "testing"

func TestSha256(t *testing.T) {
	cmdsha := "3ade5cd971524e5de1a79ab548e75c8f9f54af5110b91c6ee332d626b7572f73"
	sha, err := FileSha256Sum("tar-1.29-1-x86_64.pkg.tar.xz")
	if err != nil {
		t.Error(err)
	}

	if sha != cmdsha {
		t.Error("something went wrong they are not same\n")
	}
}
func TestSha256NonFile(t *testing.T) {
	_, err := FileSha256Sum("tar")
	if err == nil {
		t.Error(err)
	}

}
func TestVerifySha256(t *testing.T) {
	sha := "3ade5cd971524e5de1a79ab548e75c8f9f54af5110b91c6ee332d626b7572f73"
	ok, err := VerifySha256(sha, "tar-1.29-1-x86_64.pkg.tar.xz")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("not equal")
	}

}
func TestVerifySha256NoFile(t *testing.T) {
	sha := "3ade5cd971524e5de1a79ab548e75c8f9f54af5110b91c6ee332d626b7572f73"
	ok, err := VerifySha256(sha, "tar-1.29-1-x8xz")
	if err == nil {
		t.Error(err)
	}
	if ok {
		t.Error("should be false not true")
	}

}
func TestMd5(t *testing.T) {
	cmdmd5 := "e02d552239d566d7eb592d4662773ac2"
	md5, err := FileMd5Sum("tar-1.29-1-x86_64.pkg.tar.xz")
	if err != nil {
		t.Error(err)
	}

	if md5 != cmdmd5 {
		t.Error("something went wrong they are not the same\n")
	}
}

func TestMd5NoFile(t *testing.T) {
	_, err := FileMd5Sum("tarpkg.tar.xz")
	if err == nil {
		t.Error(err)
	}

}
func TestVerifyMd5(t *testing.T) {
	md5 := "e02d552239d566d7eb592d4662773ac2"
	ok, err := VerifyMd5(md5, "tar-1.29-1-x86_64.pkg.tar.xz")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("not equal")
	}
}

func TestVerifyMd5NoFile(t *testing.T) {
	md5 := "e02d552239d566d7eb592d4662773ac2"
	ok, err := VerifyMd5(md5, "tr-1.29-164.pkg.tar.xz")
	if err == nil {
		t.Error(err)
	}
	if ok {
		t.Error("should be false not true")
	}
}
