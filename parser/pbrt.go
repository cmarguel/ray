package parser

import (
	"bufio"
	"fmt"
	"os"
)

func GetDirectives(path string) []Directive {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Could not find %s\n", path)
		return []Directive{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	parser := NewParser(reader)
	parser.Parse()
	return parser.Directives
}

func concat(old1, old2 []byte) []byte {
	newslice := make([]byte, len(old1)+len(old2))
	copy(newslice, old1)
	copy(newslice[len(old1):], old2)
	return newslice
}
