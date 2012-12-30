package parser

import (
	"bufio"
)

type Symbol struct {
	tok string
	str string
}

type Lexer struct {
	reader  *bufio.Reader
	has     bool
	current byte
}

func NewLexer(reader *bufio.Reader) *Lexer {
	b, err := reader.ReadByte()
	l := &Lexer{reader, err == nil, b}
	l.skipWhitespace()
	return l
}

func (l *Lexer) hasMore() bool {
	return l.has
}

func (l *Lexer) get() {
	b, err := l.reader.ReadByte()
	l.current = b
	l.has = err == nil
}

func isSpecial(c uint8) bool {
	return c == '"' || c == '[' || c == ']' || c == '\n' || c == '#'
}

func (l *Lexer) next() Symbol {
	if !l.hasMore() {
		return Symbol{"eof", ""}
	}
	c := l.current
	if isSpecial(c) {
		l.get()
		l.skipWhitespace()
		if c == '"' {
			return Symbol{"quote", string(c)}
		} else if c == '[' {
			return Symbol{"lbr", string(c)}
		} else if c == ']' {
			return Symbol{"rbr", string(c)}
		} else if c == '\n' {
			return Symbol{"nl", string(c)}
		} else if c == '#' {
			return Symbol{"hash", string(c)}
		}
		return Symbol{"", ""}
	} else {
		acc := ""
		for l.hasMore() && !isWhitespace(l.current) && !isSpecial(l.current) {
			acc += string(l.current)
			l.get()
		}
		if l.hasMore() {
			l.skipWhitespace()
		}
		return Symbol{"ident", acc}
	}
	return Symbol{"", ""}
}

func (l *Lexer) skipWhitespace() {
	for l.hasMore() && isWhitespace(l.current) {
		l.get()
	}
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\r'
}
