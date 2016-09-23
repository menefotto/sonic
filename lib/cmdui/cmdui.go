package cmdui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sonic/lib/utils/terminal"
)

func UserConfirmation(msg string) (string, bool, error) {

	const errmsg = `Not valid input, accepted values are: 
	 YES, yes, Y, y, NO, no, N, n. \n`

	var (
		inputbool         bool
		inputstring       string
		NotValidUserInput = errors.New(errmsg)
	)

	fmt.Fprintf(os.Stdout, "%s\n", NewMsg(msg))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		inputstring := strings.TrimSpace(scanner.Text())
		switch {
		case inputstring == "YES":
			inputbool = true
			goto OK
		case inputstring == "yes":
			inputbool = true
			goto OK
		case inputstring == "Y":
			inputbool = true
			goto OK
		case inputstring == "y":
			inputbool = true
			goto OK
		case inputstring == "NO":
			inputbool = false
			goto OK
		case inputstring == "no":
			inputbool = false
			goto OK
		case inputstring == "N":
			inputbool = false
			goto OK
		case inputstring == "n":
			inputbool = false
			goto OK
		default:
			return inputstring, inputbool, NotValidUserInput
		}

	}

OK:
	return inputstring, inputbool, nil
}

func NewMsg(msg string) string {
	w, _ := terminal.GetDimensions()

	var newmsg []byte

	if len(msg) > (w - 19) {
		newmsg = []byte(msg)[:w-19]
	}
	if len(msg) < (w - 19) {
		newmsg = []byte(msg)
	}

	spacen := w - (len(msg) + 9)
	spaces := []byte(" ")

	for i := 0; i < spacen; i++ {
		spaces = append(spaces, []byte(" ")...)
	}

	return string(newmsg) + string(spaces)
}

func ProgressPrinter(msg string, tot, percent int64) {
	if tot/percent == 100 {
		fmt.Fprintf(os.Stdout, "%s%d%s\n", msg, 100, "%")
	} else {
		fmt.Fprintf(os.Stdout, "%s%d%s\r", msg, tot/percent, "%")
	}
}

type Color struct {
	list map[string]string
}

func NewColor() *Color {
	return &Color{list: map[string]string{
		"black":  "0;30",
		"blue":   "0;34",
		"green":  "0;32",
		"cyan":   "0;36",
		"red":    "0;31",
		"purple": "0;35",
		"brown":  "0;33",
		"gray":   "1;30",
		"yellow": "1;33",
		"white":  "1;37",
	},
	}
}

func (c *Color) Colorize(msg string, color string) string {
	var (
		ok    bool
		value string
	)

	value, ok = c.list[color]
	if !ok {
		return msg
	}

	return `\033[` + value + `m` + msg + `\033[0m\n`
}
