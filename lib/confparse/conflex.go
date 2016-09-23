package confparse

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type Token int

const (
	EOF Token = iota
	KEY_VALUE
	SECTION
	WHITESPACE
	NON_VALID
	COMMENT
)

const eof = rune(0)

type itemType struct {
	Tok    Token
	Values []string
}

func NewItemType(tok Token, vals ...string) *itemType {
	item := &itemType{Values: make([]string, 0), Tok: tok}
	for _, val := range vals {
		item.Values = append(item.Values, val)
	}
	return item
}

type Lexer struct {
	lex   *bytes.Buffer
	ori   []byte
	runes []rune
}

func NewLexer(r io.Reader) *Lexer {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil
	}

	return &Lexer{lex: bytes.NewBuffer(buf), ori: buf, runes: make([]rune, 0)}
}

func (l *Lexer) Scan() *itemType {
	ch := l.read()
	l.runes = append(l.runes, ch)

	switch {
	case isWhiteSpace(ch):
		l.unread()
		return l.eatWspace()
	case isSection(ch):
		return l.eatSection()
	case isValue(ch):
		return l.eatKeyValue()
	case isComment(ch):
		l.unread()
		return l.eatComment()
	}

	if ch == eof {
		return NewItemType(EOF, "")
	}

	return NewItemType(NON_VALID, string(ch))
}

func (l *Lexer) read() rune {
	ch, _, err := l.lex.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (l *Lexer) unread() { l.lex.UnreadRune() }

func (l *Lexer) eatWspace() *itemType {

	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == eof {
			break
		}
		if !isWhiteSpace(ch) {
			l.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NewItemType(WHITESPACE, buf.String())
}

func (l *Lexer) eatKeyValue() *itemType {
	var value bytes.Buffer
	value.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == eof || ch == '\n' {
			break
		}
		if isDigit(ch) || isLetter(ch) || isValid(ch) {
			value.WriteRune(ch)
		}

	}

	for i := len(l.runes) - 1; i >= 0; i-- {
		if l.runes[i] == '\n' {
			l.runes = l.runes[i+1:]
			break
		}
	}

	var (
		count, index int
		key          string
	)

	count = strings.Count(string(l.runes), "=")
	if count == 1 {
		index = strings.Index(string(l.runes), "=")
		key = string(l.runes[:len(l.runes)-1])
	}
	if count > 1 {
		index = strings.LastIndex(string(l.runes[:len(l.runes)-1]), "=")
		key = string(l.runes[index:])
	}

	return NewItemType(KEY_VALUE,
		strings.TrimSpace(strings.Trim(key, "=")),
		strings.TrimSpace(value.String()))
}

func (l *Lexer) eatComment() *itemType {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == eof || ch == '\n' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NewItemType(COMMENT, strings.TrimSpace(buf.String()))
}

func (l *Lexer) eatSection() *itemType {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == eof {
			break
		}
		if ch == ']' {
			l.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return NewItemType(SECTION, strings.TrimSpace(buf.String()))
}

func (l *Lexer) findLine(word string) (int, error) {
	copy := bytes.NewBuffer(l.ori)
	if copy == nil {
		return -1, fmt.Errorf("can't allocate slice\n")
	}
	regex, err := regexp.Compile(word)
	if err != nil {
		return -1, err
	}

	line := 0
	for {
		str, err := copy.ReadString('\n')
		if err != nil {
			return line, err
		}
		if ok := regex.Match([]byte(str)); ok {
			return line, err
		}
		line++
	}

}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isComment(ch rune) bool {
	return ch == ';' || ch == '#'
}

func isValid(ch rune) bool {
	return ch == '.' || ch == '@' || ch == '/' || ch == ',' || ch == '-' || ch == ':' || ch == '_'
}

func isValue(ch rune) bool {
	return ch == '='
}

func isSection(ch rune) bool {
	return ch == '['
}
