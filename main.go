package main

import (
	"fmt"
	"github.com/gookit/slog"
	"lox_go/lox"
	"os"
)

func main() {

	slog.SetLogLevel(slog.InfoLevel)

	args := os.Args
	if len(args) > 2 {
		fmt.Printf("Usage: %s [script]\n", args[0])
		os.Exit(64)
	} else if len(args) == 2 {
		lox.RunFile(args[1])
	} else {
		lox.RunPrompt()
	}
}
