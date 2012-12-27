package parser

import (
	"strings"
)

/* line = comment | spec
 * spec = ident {arg}
 * arg = ident | "{ident}" array
 * array = [{ident}]
 */

type Directive struct {
	Name string
	Args []string
}

type Parser struct {
	str       string
	sym       Symbol
	lex       *Lexer
	Directive Directive
}

func NewParser(s string) *Parser {
	l := NewLexer(s)

	return &Parser{s, l.next(), l, Directive{}}
}

func (p *Parser) Parse() bool {
	if p.sym.str[0] == '#' {
		return false
	}
	p.spec()
	return true
}

func (p *Parser) spec() {
	p.Directive.Name = p.sym.str
	p.accept("ident")
	for p.lex.hasMore() && (p.sym.tok == "ident" || p.sym.tok == "quote") {
		p.arg()
	}
}

func (p *Parser) add(s string) {
	p.Directive.Args = append(p.Directive.Args, s)
}

func (p *Parser) arg() {
	if p.sym.tok == "ident" {
		p.add(p.sym.str)
		p.accept("ident")
	} else if p.accept("quote") {
		id := ""
		id += p.sym.str + " "
		for p.accept("ident") {
			if p.sym.tok == "ident" {
				id += p.sym.str + " "
			}
		}
		p.expect("quote")
		p.add("\"" + strings.TrimSpace(id) + "\"")

		p.array()
	} else {
		panic("parse error")
	}
}

func (p *Parser) array() {
	p.accept("lbr")
	id := ""
	id += p.sym.str + " "
	for p.accept("ident") {
		if p.sym.tok == "ident" {
			id += p.sym.str + " "
		}
	}
	id = strings.TrimSpace(id)
	p.expect("rbr")
	p.add(id)
}

func (p *Parser) next() {
	if p.lex.hasMore() {
		p.sym = p.lex.next()
	}
}

func (p *Parser) accept(s string) bool {
	if p.sym.tok == s {
		p.next()
		return true
	}
	return false
}

func (p *Parser) expect(s string) bool {
	if p.accept(s) {
		return true
	}
	panic("Parse error!")
	return false
}
