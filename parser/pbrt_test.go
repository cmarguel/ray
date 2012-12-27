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
	fmt.Printf("|%s|%s|%s|%s|", t1, t2, t3, t4)
	if t1 != "foobar" || t2 != "[" || t3 != "1" || t4 != "2" || t5 != "3" || t6 != "]" {
		t.Fail()
	}
}
