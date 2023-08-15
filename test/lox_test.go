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

func TestLox9Flow(t *testing.T) {
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

func TestLox9IfStmt(t *testing.T) {
	lox.Eval(codeIfStmt)
}

const codeFunction = `
fun count(n) {
  if (n > 1) count(n - 1);
  print n;
}

count(3);

fun add(a, b) {
  return function() {
    return a + b;
  };
}
var addFunc = add(1 + 2);
print addFunc(); // "<fn add>".

`

func TestLox10Function(t *testing.T) {
	lox.Eval(codeFunction)
}

const codeFunctionFib = `
fun fib(n) {
  if (n <= 1) return n;
  return fib(n - 2) + fib(n - 1);
}

var before = clock();
for (var i = 0; i < 20; i = i + 1) {
  print fib(i) + "\n";
}
var after = clock();
print after - before;
`

func TestLox10FunctionFib(t *testing.T) {
	lox.Eval(codeFunctionFib)
}

const codeFunctionClosure = `
fun makeCounter() {
  var i = 0;
  fun count() {
    i = i + 1;
    print i;
  }

  return count;
}

var counter = makeCounter();
counter(); // "1".
counter(); // "2".
counter = makeCounter();
counter(); // "1".
counter(); // "2".
`

func TestLox10FunctionClosure(t *testing.T) {
	lox.Eval(codeFunctionClosure)
}

const code11Binding = `
var a = "global";
{
  fun showA() {
    print a;
  }

  showA();
  var a = "block";
  showA();
}
`

func TestLox11Resolving(t *testing.T) {
	lox.Eval(code11Binding)
}
