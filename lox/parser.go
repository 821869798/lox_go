package lox

type Parser struct {
	tokens  []*Token
	current int
}

func NewParse(tokens []*Token) *Parser {
	p := &Parser{
		tokens:  tokens,
		current: 0,
	}
	return p
}

func (p *Parser) parse() []Stmt {
	defer func() {
		if err := recover(); err != nil {
			//v, ok := err.(*RuntimeError)
			//if ok {
			//	reportRuntimeError(v)
			//} else {
			//	panic(err)
			//}
		}
	}()
	statements := make([]Stmt, 0, 4)
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) declaration() Stmt {
	if p.match(TokenType_VAR) {
		return p.varDeclaration()
	}

	return p.statement()
}

func (p *Parser) statement() Stmt {
	if p.match(TokenType_PRINT) {
		return p.printStatement()
	}
	if p.match(TokenType_LEFT_BRACE) {
		return NewBlock(p.block())
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() Stmt {
	value := p.expression()
	p.consume(TokenType_SEMICOLON, "Expect ';' after value.")
	return NewPrint(value)
}

func (p *Parser) varDeclaration() Stmt {
	name := p.consume(TokenType_IDENTIFIER, "Expect variable name.")

	var initializer Expr = nil
	if p.match(TokenType_EQUAL) {
		initializer = p.expression()
	}
	p.consume(TokenType_SEMICOLON, "Expect ';' after variable declaration.")

	return NewVarStmt(name, initializer)
}

func (p *Parser) expressionStatement() Stmt {
	expr := p.expression()
	p.consume(TokenType_SEMICOLON, "Expect ';' after expression.")
	return NewExpression(expr)
}

func (p *Parser) block() []Stmt {
	statements := make([]Stmt, 0, 4)
	for !p.check(TokenType_RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}

	p.consume(TokenType_RIGHT_BRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) assignment() Expr {
	expr := p.equality()
	if p.match(TokenType_EQUAL) {
		equals := p.previous()
		value := p.assignment()

		variable, ok := expr.(*Variable)
		if ok {
			name := variable.name
			return NewAssign(name, value)
		}

		reportErrorToken(equals, "Invalid assignment target.")
	}

	return expr
}

func (p *Parser) expression() Expr {
	return p.assignment()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(TokenType_BANG_EQUAL, TokenType_EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(TokenType_GREATER, TokenType_GREATER_EQUAL, TokenType_LESS, TokenType_LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) match(types ...TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().tokenType == tokenType
}

func (p *Parser) advance() *Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().tokenType == TokenType_EOF
}

func (p *Parser) peek() *Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *Token {
	return p.tokens[p.current-1]
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(TokenType_MINUS, TokenType_PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(TokenType_SLASH, TokenType_STAR) {
		operator := p.previous()
		right := p.unary()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(TokenType_BANG, TokenType_MINUS) {
		operator := p.previous()
		right := p.unary()
		return NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(TokenType_FALSE) {
		return NewLiteral(false)
	}
	if p.match(TokenType_TRUE) {
		return NewLiteral(true)
	}
	if p.match(TokenType_NIL) {
		return NewLiteral(nil)
	}

	if p.match(TokenType_NUMBER, TokenType_STRING) {
		return NewLiteral(p.previous().literal)
	}

	if p.match(TokenType_IDENTIFIER) {
		return NewVariable(p.previous())
	}

	if p.match(TokenType_LEFT_PAREN) {
		expr := p.expression()
		p.consume(TokenType_RIGHT_PAREN, "Expect ')' after expression.")
		return NewGrouping(expr)
	}

	message := "Expect expression."
	reportErrorToken(p.peek(), message)
	panic(message)
}

func (p *Parser) consume(tokenType TokenType, message string) *Token {
	if p.check(tokenType) {
		return p.advance()
	}

	reportErrorToken(p.peek(), message)
	panic(message)
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().tokenType == TokenType_SEMICOLON {
			return
		}

		switch p.peek().tokenType {
		case TokenType_CLASS, TokenType_FUN, TokenType_VAR, TokenType_FOR, TokenType_IF, TokenType_WHILE, TokenType_PRINT, TokenType_RETURN:
			return
		}

		p.advance()
	}
}
