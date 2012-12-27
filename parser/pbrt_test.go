package parser

import (
	"fmt"
	"testing"
)

func Test_lex(t *testing.T) {
	l := NewLexer("foobar [1 2 3]")
	t1 := l.next()
	t2 := l.next()
	t3 := l.next()
	t4 := l.next()
	t5 := l.next()
	t6 := l.next()
	fmt.Printf("|%s|%s|%s|%s|\n", t1, t2, t3, t4)
	if t1.str != "foobar" || t2.str != "[" || t3.str != "1" || t4.str != "2" || t5.str != "3" || t6.str != "]" {
		t.Fail()
	}
}

func Test_comment(t *testing.T) {
	p := NewParser("#foobar [1 2 3]")
	if p.Parse() {
		t.Fail()
	}
}

func Test_no_args(t *testing.T) {
	p := NewParser("foobar")
	if !p.Parse() {
		t.Fail()
	}
	if p.Directive.Name != "foobar" || len(p.Directive.Args) > 0 {
		t.Fail()
	}
}

func Test_num_args(t *testing.T) {
	p := NewParser("foobar 1 22 33")
	if !p.Parse() {
		t.Fail()
	}
	if p.Directive.Name != "foobar" || len(p.Directive.Args) != 3 {
		fmt.Printf("len was %d, first arg was |%s| and second was |%s|\n", len(p.Directive.Args), p.Directive.Args[0], p.Directive.Args[1])
		t.Fail()
	}
	fmt.Println(p.Directive.Args)
}

func Test_named_array(t *testing.T) {
	p := NewParser("foobar \"baz qux\" [1 2 3]")
	if !p.Parse() {
		t.Fail()
	}
	if p.Directive.Name != "foobar" || len(p.Directive.Args) != 2 {
		t.Fail()
	}
	if p.Directive.Args[0] != "\"baz qux\"" || p.Directive.Args[1] != "1 2 3" {
		fmt.Printf("first arg was |%s| and second was |%s|\n", p.Directive.Args[0], p.Directive.Args[1])
		t.Fail()
	}
	fmt.Println(p.Directive.Args)

}

func Test_singleton_strings(t *testing.T) {
	p := NewParser("foobar \"alpha\" \"beta\"")
	if !p.Parse() {
		t.Fail()
	}
	if p.Directive.Name != "foobar" || len(p.Directive.Args) != 2 {
		t.Fail()
	}
	if p.Directive.Args[0] != "\"alpha\"" || p.Directive.Args[1] != "\"beta\"" {
		t.Fail()
	}
}
