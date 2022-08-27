package main

import (
	"os"
	"strings"

	"github.com/ryicoh/mylang"
)

func main() {
	lex := mylang.NewLexer(strings.NewReader(os.Args[1]))
	prog, err := mylang.Parse(lex)
	if err != nil {
		panic(err)
	}

	asm := mylang.Codegen(prog)
	os.WriteFile("a.ll", []byte(asm), 0600)
}
