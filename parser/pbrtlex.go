package parser

type Symbol struct {
	tok string
	str string
}

type Lexer struct {
	str string
	pos int
}

func NewLexer(s string) *Lexer {
	l := &Lexer{s, 0}
	l.skipWhitespace()
	return l
}

func (l *Lexer) hasMore() bool {
	return l.pos <= len(l.str)
}

func (l *Lexer) get() {
	if l.hasMore() {
		l.pos++
	}
}

func isSpecial(c uint8) bool {
	return c == '"' || c == '[' || c == ']'
}

func (l *Lexer) next() Symbol {
	if l.pos == len(l.str) {
		return Symbol{"eof", ""}
	}
	c := l.str[l.pos]
	if isSpecial(c) {
		l.get()
		l.skipWhitespace()
		if c == '"' {
			return Symbol{"quote", string(c)}
		} else if c == '[' {
			return Symbol{"lbr", string(c)}
		} else if c == ']' {
			return Symbol{"rbr", string(c)}
		}
		return Symbol{"", ""}
	} else {
		acc := ""
		for l.pos < len(l.str) && l.str[l.pos] != ' ' && !isSpecial(l.str[l.pos]) {
			acc += string(l.str[l.pos])
			l.get()
		}
		l.skipWhitespace()
		return Symbol{"ident", acc}
	}
	return Symbol{"", ""}
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.str) && l.str[l.pos] == ' ' {
		l.get()
	}
}
