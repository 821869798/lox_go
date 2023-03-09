package lox

import (
	"bufio"
	"fmt"
	log "github.com/FishGoddess/logit"
	"io/ioutil"
	"os"
	"strings"
)

var hadError = false
var hadRuntimeError = false

var interpreter = NewInterpreter()

func run(source string) {

	scanner := NewScanner(source)
	tokens := scanner.scanTokens()

	parser := NewParse(tokens)
	statements := parser.parse()
	if hadError {
		return
	}

	interpreter.interpret(statements)
}

func RunFile(filename string) {
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Info("Error reading file: %s", filename)
		return
	}

	run(string(code))

	if hadError {
		os.Exit(65)
	}

	if hadRuntimeError {
		os.Exit(70)
	}
}

func RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Error("input error:%v", err)
			break
		}
		code := strings.TrimRight(input, "\n")
		run(code)
		hadError = false
		fmt.Print("\n")
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
	log.Error("<error>[line %d] Error%s: %s", line, where, message)
	hadError = true
}

func reportRuntimeError(err *RuntimeError) {
	log.Error(err.Error())
	hadRuntimeError = true
}
