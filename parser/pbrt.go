package parser

/* line = comment | spec
 * spec = directive {arg}
 * arg = num | quoted | array
 * quoted = "{_}"
 * array = [{_}]
 */

type Parser struct {
	str string
	sym byte
	pos int
}

func NewParser(s string) *Parser {
	return &Parser{s, s[0], 0}
}

func (p *Parser) Parse() {
	p.spec()
}

func (p *Parser) spec() {

}

func (p *Parser) next() {
	p.pos++
}

func (p *Parser) accept(b byte) bool {
	if p.sym == b {
		p.next()
		return true
	}
	return false
}

func (p *Parser) expect(b byte) bool {
	if p.accept(b) {
		return true
	}
	panic("Parse error!")
	return false
}
