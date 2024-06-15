package main

import (

	"github.com/sef-computin/bmstu-cc2024-lab3/grammar"
)

func main() {
	input := "begin a = true & ~false; begin c = true; c = a & c end; b = a ! false; end"
	// input := "begin a = true & ~false end"
	lexer := grammar.NewLexer(input)
	parser := grammar.NewParser(lexer)

	parser.ParseProgram()
	// fmt.Println("Parsing completed successfully")

	grammar.DrawTree("grammarlab3", parser.ParseTree)
}
