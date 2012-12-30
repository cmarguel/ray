package parser

import (
	"bufio"
	"fmt"
	"strings"
)

/* file = {directive} 
 * directive = 
 *     # ... nl 
 *     | spec nl
 * spec = ident {arg}
 * arg = 
 *     ident 
 *     | quote {ident} quote [lbr array rbr]
 * array = {[nl] ident [nl]}
 */

type Directive struct {
	Name string
	Args []string
}

type Parser struct {
	sym        Symbol
	lex        *Lexer
	Directives []Directive
}

func (p Parser) curr() *Directive {
	return &p.Directives[len(p.Directives)-1]
}

func NewParser(reader *bufio.Reader) *Parser {
	l := NewLexer(reader)

	return &Parser{l.next(), l, make([]Directive, 0)}
}

func (p *Parser) Parse() bool {
	for !p.accept("eof") {
		p.line()
	}
	return true
}

func (p *Parser) line() {
	if p.accept("nl") || p.accept("eof") {

	} else if p.accept("hash") {
		for !p.accept("nl") && !p.accept("eof") {
			p.next()
		}
	} else {
		dir := *new(Directive)
		fmt.Println("creating directive for ", p.sym)
		p.Directives = append(p.Directives, dir)
		p.spec()
	}
}

func (p *Parser) spec() {
	p.curr().Name = p.sym.str
	p.accept("ident")
	for p.sym.tok == "ident" || p.sym.tok == "quote" || p.sym.tok == "lbr" {
		p.arg()
	}
}

func (p *Parser) add(s string) {
	p.curr().Args = append(p.curr().Args, s)
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

		if p.sym.tok == "lbr" {
			p.array()
		}
	} else if p.sym.tok == "lbr" {
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
	p.sym = p.lex.next()
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
