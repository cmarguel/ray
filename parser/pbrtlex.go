package parser

type Lexer struct {
	str string
	pos int
}

func NewLexer(s string) *Lexer {
	return &Lexer{s, 0}
}

func (l *Lexer) hasMore() bool {
	return l.pos < len(l.str)
}

func (l *Lexer) get() {
	if l.hasMore() {
		l.pos++
	}
}

func isSpecial(c uint8) bool {
	return c == '"' || c == '[' || c == ']'
}

func (l *Lexer) next() string {
	c := l.str[l.pos]
	if isSpecial(c) {
		l.get()
		l.skipWhitespace()
		return string(c)
	} else {
		acc := ""
		for l.str[l.pos] != ' ' && !isSpecial(l.str[l.pos]) {
			acc += string(l.str[l.pos])
			l.get()
		}
		l.skipWhitespace()
		return acc
	}
	return "!!!"
}

func (l *Lexer) skipWhitespace() {
	for l.hasMore() && l.str[l.pos] == ' ' {
		l.get()
	}
}
