package test

import (
	"lox_go/lox"
	"testing"
)

const codeFlow string = `
var a = 0;
var temp;

for (var b = 1; a < 10000; b = temp + b) {
  print a + "\n";
  temp = a;
  a = b;
}
`

func TestLoxFlow(t *testing.T) {
	lox.Eval(codeFlow)
}

const codeIfStmt = `
var condition = true;
if (condition) {
  print "yes\n";
} else {
  print "no";
}
`

func TestLoxIfStmt(t *testing.T) {
	lox.Eval(codeIfStmt)
}
