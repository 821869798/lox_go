package lox

import (
	"bufio"
	"fmt"
	"github.com/gookit/slog"
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

	resolver := NewResolver(interpreter)
	resolver.resolveStmt(statements)

	interpreter.interpret(statements)
}

func Eval(code string) {
	run(code)
}

func RunFile(filename string) {
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		slog.Infof("Error reading file: %s", filename)
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
			slog.Errorf("input error:%v", err)
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
	slog.Errorf("<error>[line %d] Error%s: %s", line, where, message)
	hadError = true
}

func reportRuntimeError(err *RuntimeError) {
	slog.Errorf(err.Error())
	hadRuntimeError = true
}
