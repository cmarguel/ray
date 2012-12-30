package parser

import (
	"bufio"
	"fmt"
	"os"
	"ray/geom"
	"strings"
	"testing"
)

func newStringReader(s string) *bufio.Reader {
	b := bufio.NewReader(strings.NewReader(s))
	return b
}

func Test_lex(t *testing.T) {
	l := NewLexer(newStringReader("foobar [1 2 3]"))
	t1 := l.next()
	t2 := l.next()
	t3 := l.next()
	t4 := l.next()
	t5 := l.next()
	t6 := l.next()
	fmt.Printf("|%s|%s|%s|%s|\n", t1, t2, t3, t4)
	if l.hasMore() {
		fmt.Println("Lexer should have nothing left")
		t.Fail()
	}
	if t1.str != "foobar" || t2.str != "[" || t3.str != "1" || t4.str != "2" || t5.str != "3" || t6.str != "]" {
		t.Fail()
	}
}

func Test_lex_with_newlines(t *testing.T) {
	l := NewLexer(newStringReader("foobar [1 2 3]\nbarbaz 1"))
	t1 := l.next()
	l.next()
	l.next()
	l.next()
	t5 := l.next()
	t6 := l.next()
	t7 := l.next()
	t8 := l.next()
	t9 := l.next()
	t10 := l.next()
	fmt.Printf("|%s|%s|%s|%s|%s|\n", t5, t6, t7, t8, t9)
	if l.hasMore() {
		fmt.Println("Lexer should have nothing left")
		t.Fail()
	}
	if t1.str != "foobar" || t5.str != "3" || t6.str != "]" || t7.tok != "nl" || t8.str != "barbaz" || t9.str != "1" {
		fmt.Println(l.current)
		t.Fail()
	}
	if t10.tok != "eof" {
		fmt.Println("EOF not reached")
		t.Fail()
	}
}

func Test_comment(t *testing.T) {
	p := NewParser(newStringReader("#foobar [1 2 3]"))
	p.Parse()
	if len(p.Directives) > 0 {
		fmt.Println(p.Directives)
		t.Fail()
	}
}

func Test_no_args(t *testing.T) {
	p := NewParser(newStringReader("foobar"))
	p.Parse()
	if len(p.Directives) > 1 {
		t.Fail()
	}
	if p.curr().Name != "foobar" || len(p.curr().Args) > 0 {
		fmt.Println(p.curr().Args)
		t.Fail()
	}
}

func Test_num_args(t *testing.T) {
	p := NewParser(newStringReader("foobar 1 22 33"))
	if !p.Parse() {
		t.Fail()
	}
	if p.curr().Name != "foobar" || len(p.curr().Args) != 3 {
		fmt.Println("number of directives: ", len(p.Directives))
		fmt.Printf("len was %d\n", len(p.curr().Args))
		fmt.Printf("len was %d, first arg was |%s| and second was |%s|\n", len(p.curr().Args), p.curr().Args[0], p.curr().Args[1])
		t.Fail()
	}
}

func Test_named_array(t *testing.T) {
	p := NewParser(newStringReader("foobar \"baz qux\" [1 2 3]"))
	if !p.Parse() {
		t.Fail()
	}
	if p.curr().Name != "foobar" || len(p.curr().Args) != 2 {
		t.Fail()
	}
	if p.curr().Args[0] != "\"baz qux\"" || p.curr().Args[1] != "1 2 3" {
		fmt.Printf("first arg was |%s| and second was |%s|\n", p.curr().Args[0], p.curr().Args[1])
		t.Fail()
	}
	fmt.Println(p.curr().Args)

}

func Test_singleton_strings(t *testing.T) {
	p := NewParser(newStringReader("foobar \"alpha\" \"beta\""))
	if !p.Parse() {
		t.Fail()
	}
	if p.curr().Name != "foobar" || len(p.curr().Args) != 2 {
		t.Fail()
	}
	if p.curr().Args[0] != "\"alpha\"" || p.curr().Args[1] != "\"beta\"" {
		t.Fail()
	}
}

func Dont_Test_getDirectives(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Printf("working from %s\n", dir)
	d := GetDirectives("../scenes/cornell-mlt.pbrt")
	fmt.Printf("Loaded %d curr()s\n", len(d))
	if len(d) != 39 {
		t.Fail()
	}
}

func Test_config(t *testing.T) {
	conf := LoadConfig("../scenes/cornell-mlt.pbrt")
	if conf.Translate.Distance(geom.NewVector3(-278., -273, 500)) > 0.005 {
		fmt.Printf("Translate was way off. Actual values: %s", conf.Translate)
		t.Fail()
	}
	if conf.Fov != 55. {
		fmt.Printf("FOV was %f\n", conf.Fov)
		t.Fail()
	}
}
