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

const codeFunction2 = `
fun Add(a,b) {
	return a+b;
}

fun Hello(num) {
	print "Hello " + num;
}

Hello(Add(1,2));
`

func TestLox10Function2(t *testing.T) {
	lox.Eval(codeFunction2)
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
  return a;
  fun showA() {
    print a;
  }

  showA();
  a = "block";
  showA();
}
`

func TestLox11Resolving(t *testing.T) {
	lox.Eval(code11Binding)
}

const code12Class1 = `
class Bagel {}
var bagel = Bagel();
print bagel; // Prints "Bagel instance".
`

func TestLox12Class1(t *testing.T) {
	lox.Eval(code12Class1)
}

const code12Class2 = `
class Bacon {
  eat() {
    print "Crunch crunch crunch!";
  }
}

Bacon().eat(); // Prints "Crunch crunch crunch!".
`

func TestLox12Class2(t *testing.T) {
	lox.Eval(code12Class2)
}

const code12Class3 = `
class Person {
  init() {
	this.name = "123";
  }
  sayName() {
    print this.name;
  }
}

var jane = Person();
print jane.name + "\n";
jane.name = "Jane";

var bill = Person();
bill.name = "Bill";

bill.sayName = jane.sayName;
bill.sayName(); // ?
`

func TestLox12Class3(t *testing.T) {
	lox.Eval(code12Class3)
}

const code13Inheritance1 = `
class Doughnut {
  cook() {
    print "Fry until golden brown.";
  }
}

class BostonCream < Doughnut {}

BostonCream().cook();
`

func TestLox13Inheritance1(t *testing.T) {
	lox.Eval(code13Inheritance1)
}

const code13Inheritance2 = `
class A {
  method() {
    print "A method";
  }
}

class B < A {
  method() {
    print "B method";
  }

  test() {
    super.method();
  }
}

class C < B {}

C().test();
`

func TestLox13Inheritance2(t *testing.T) {
	lox.Eval(code13Inheritance2)
}

const code13Inheritance3 = `
class Doughnut {
  init() {
    this.name = "123";
	print "Doughnut init\n";
  }
  cook() {
    print "Fry until golden brown.";
  }
}

class BostonCream < Doughnut {
  init() {
    super.init();
	print "BostonCream init\n";
	print this.name;
  }
}

BostonCream().cook();
`

func TestLox13Inheritance3(t *testing.T) {
	lox.Eval(code13Inheritance3)
}
