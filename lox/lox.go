package lox

import (
	"fmt"
	"io/ioutil"
	"os"
)

var hadError = false

func run(source string) {
	fmt.Println("run:" + source)

	scanner := NewScanner(source)
	tokens := scanner.scanTokens()

	parser := NewParse(tokens)
	expression := parser.parse()
	if hadError {
		return
	}

	NewAstPrinter().print(expression)
}

func RunFile(filename string) {
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", filename)
		return
	}
	run(string(code))

	if hadError {
		os.Exit(65)
	}
}

func RunPrompt() {
	for {
		fmt.Print("> ")
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			fmt.Println(err)
			break
		}
		run(line)
		hadError = false
	}
}

func reportError(line int, message string) {
	report(line, "", message)
}

func reportErrorToken(token *Token, message string) {
	if token.tokenType == TokenType_EOF {
		report(token.line, " at end", message)
	} else {
		report(token.line, " at '"+token.lexeme+"'", message)
	}
}

func report(line int, where string, message string) {
	fmt.Printf("<error>[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}
