package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sef-computin/bmstu-cc2024-lab3/grammar"
)

func main() {

	texts := []string{
		"begin axiom = true end",
		"begin lie = false; nlie = ~lie end",
		"begin a = true & ~false end",
		"begin a = true & ~false; begin c = true; c = a & c end; b = a ! false end",
		"begin begin a = true end end",
		"begin b = false; a = ~true ! true ! ~b & false ! b end",
	}

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Неправильное число аргументов")
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if num > len(texts)-1 || num < 0 {
		num = 0
	}

	test(texts[num])
}

func test(input string) {
	lexer := grammar.NewLexer(input)
	parser := grammar.NewParser(lexer)

	parser.ParseProgram()
	fmt.Println("Parsing completed successfully")

	grammar.DrawTree("grammarlab3", parser.ParseTree)
}
