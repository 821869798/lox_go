package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//生成expr.go文件

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s <output directory>\n", args[0])
		os.Exit(64)
	}

	outputDir := args[1]
	defineAst(outputDir, "Expr", []string{
		"AssignExpr   : name *Token, value Expr",
		"BinaryExpr   : left Expr, operator *Token, right Expr",
		"CallExpr     : callee Expr, paren *Token, arguments []Expr",
		"GetExpr      : object Expr, name *Token",
		"GroupingExpr : expression Expr",
		"LiteralExpr  : value interface{}",
		"LogicalExpr  : left Expr, operator *Token, right Expr",
		"SetExpr	  : object Expr, name *Token, value Expr",
		"SuperExpr	  : keyword *Token, method *Token",
		"ThisExpr     : keyword *Token",
		"UnaryExpr    : operator *Token, right Expr",
		"VariableExpr : name *Token",
	})

	defineAst(outputDir, "Stmt", []string{
		"BlockStmt      : statements []Stmt",
		"ClassStmt      : name *Token, superclass *VariableExpr, methods []*FunctionStmt",
		"ExpressionStmt : expression Expr",
		"FunctionStmt   : name *Token, params []*Token, body []Stmt",
		"IfStmt         : condition Expr, thenBranch Stmt," +
			" elseBranch Stmt",
		"PrintStmt      : expression Expr",
		"ReturnStmt     : keyword *Token, value Expr",
		"VarStmt    : name *Token, initializer Expr",
		"WhileStmt  : condition Expr, body Stmt",
	})
}

func defineAst(ouputDir string, baseName string, exprTypes []string) {
	path := filepath.Join(ouputDir, strings.ToLower(baseName)+".go")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.WriteString("package lox\n\n")
	f.WriteString(fmt.Sprintf("type %s interface {\n}\n\n", baseName))
	// 新增部分开始
	// The AST classes.
	for _, exprType := range exprTypes {
		exprStrs := strings.Split(exprType, ":")
		className := strings.TrimSpace(exprStrs[0])
		fields := strings.TrimSpace(exprStrs[1])
		defineType(f, baseName, className, fields)
	}

	defineVisitor(f, baseName, exprTypes)
}

func defineType(f *os.File, baseName string, className string, fieldList string) {
	f.WriteString(fmt.Sprintf("type %s struct{\n", className))

	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		f.WriteString("\t" + field + "\n")
	}
	f.WriteString("}\n\n")

	f.WriteString(fmt.Sprintf("func New%s(%s)*%s{\n", className, fieldList, className))
	classVar := strings.ToLower(className[:1])
	f.WriteString(fmt.Sprintf("\t%s := &%s{\n", classVar, className))
	for _, field := range fields {
		name := strings.Split(field, " ")[0]

		f.WriteString(fmt.Sprintf("\t\t%s: %s,\n", name, name))
	}
	f.WriteString("\t}\n")
	f.WriteString(fmt.Sprintf("\treturn %s\n", classVar))
	f.WriteString("}\n\n")
}

func defineVisitor(f *os.File, baseName string, exprTypes []string) {
	f.WriteString(fmt.Sprintf("type %sVisitor interface{\n", baseName))

	for _, exprType := range exprTypes {
		exprStrs := strings.Split(exprType, ":")
		typeName := strings.TrimSpace(exprStrs[0])
		f.WriteString(fmt.Sprintf("\tVisit%s(%s *%s)\n", typeName, strings.ToLower(typeName), typeName))
	}

	f.WriteString("}\n\n")
	varName := strings.ToLower(baseName)[:1]
	f.WriteString(fmt.Sprintf("func Visitor%s(v %sVisitor,%s %s){\n", baseName, baseName, varName, baseName))
	f.WriteString(fmt.Sprintf("\tswitch %s.(type){\n", varName))

	for _, exprType := range exprTypes {
		exprStrs := strings.Split(exprType, ":")
		typeName := strings.TrimSpace(exprStrs[0])
		f.WriteString(fmt.Sprintf("\tcase *%s:\n", typeName))
		f.WriteString(fmt.Sprintf("\t\tv.Visit%s(%s.(*%s))\n", typeName, varName, typeName))
	}
	f.WriteString("\t}\n")

	f.WriteString("}\n\n")

	//添加带泛型的Visitor
	f.WriteString(fmt.Sprintf("type %sVisitorWithVal[T any] interface{\n", baseName))

	for _, exprType := range exprTypes {
		exprStrs := strings.Split(exprType, ":")
		typeName := strings.TrimSpace(exprStrs[0])
		f.WriteString(fmt.Sprintf("\tVisit%s(%s *%s) T\n", typeName, strings.ToLower(typeName), typeName))
	}

	f.WriteString("}\n\n")
	f.WriteString(fmt.Sprintf("func Visitor%sWithVal[T any](v %sVisitorWithVal[T],%s %s) T{\n", baseName, baseName, varName, baseName))
	f.WriteString(fmt.Sprintf("\tswitch %s.(type){\n", varName))

	for _, exprType := range exprTypes {
		exprStrs := strings.Split(exprType, ":")
		typeName := strings.TrimSpace(exprStrs[0])
		f.WriteString(fmt.Sprintf("\tcase *%s:\n", typeName))
		f.WriteString(fmt.Sprintf("\t\treturn v.Visit%s(%s.(*%s))\n", typeName, varName, typeName))
	}
	f.WriteString(fmt.Sprintf("\tdefault:\n\t\tpanic(\"can't find %s\")\n", baseName))
	f.WriteString("\t}\n")

	f.WriteString("}\n\n")

}
