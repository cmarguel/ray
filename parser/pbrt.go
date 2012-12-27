package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetDirectives(path string) []Directive {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Could not find %s\n", path)
		return []Directive{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	directives := make([]Directive, 0)
	for {
		part, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		for prefix {
			var p []byte
			p, prefix, _ = reader.ReadLine()
			part = concat(part, p)
		}
		// fmt.Println("|", string(part), "|")
		line := strings.TrimSpace(string(part))
		if len(line) > 0 {
			parser := NewParser(line)

			if parser.Parse() {
				directives = append(directives, parser.Directive)
			}
		}
	}

	return directives
}

func concat(old1, old2 []byte) []byte {
	newslice := make([]byte, len(old1)+len(old2))
	copy(newslice, old1)
	copy(newslice[len(old1):], old2)
	return newslice
}
