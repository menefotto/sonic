package cmdui

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sonic/lib/utils/terminal"
)

func TestSetUp(t *testing.T) {
	f, err := os.Create("out")
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = f
}

func TestCmdUiOk(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("yes\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}
func TestCmdUiOk2(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("YES\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}
func TestCmdUiOk3(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("y\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}
func TestCmdUiOk7(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("Y\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCmdUiOk4(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("no\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}
func TestCmdUiOk5(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("n\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCmdUiOk6(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("N\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCmdUiOk8(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("NO\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err != nil {
		t.Fatal(err)
	}

}
func TestCmdUiNotOk(t *testing.T) {
	in, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = in.WriteString("aya\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = in

	msg := "Are you sure you want to proceed [yes/no]:"
	_, _, err = UserConfirmation(msg)
	if err == nil {
		t.Fatal(err)
	}

}

func TestTearDown(t *testing.T) {
	os.Remove("out")
}

func TestNewMsg(t *testing.T) {
	msg := "fal;hkfl;ashdfl;hasdl;fhl;asdhfl;hsadl;kfhlksadhlkfsdahl;fkla;sflk;sdhl;khjfdafkljadfjadfadljflaksjfldksa"
	cutmsg := NewMsg(msg)
	w, _ := terminal.GetDimensions()
	if len(cutmsg) > w {
		t.Error("msg string has not been cut to fit terminal")
	}

	NewMsg("ciao")
	ProgressPrinter(cutmsg, 1, 3)
	ProgressPrinter("done", 10000, 100)
}

func TestColor(t *testing.T) {
	f, err := os.Create("out")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	os.Stdout = f
	color := NewColor()

	fmt.Println(color.Colorize("hello color!", "green"))
}
