package lox

type TokenType int

const (
	TokenType_None TokenType = iota

	// Single-character tokens.
	TokenType_LEFT_PAREN
	TokenType_RIGHT_PAREN
	TokenType_LEFT_BRACE
	TokenType_RIGHT_BRACE
	TokenType_COMMA
	TokenType_DOT
	TokenType_MINUS
	TokenType_PLUS
	TokenType_SEMICOLON
	TokenType_SLASH
	TokenType_STAR

	// One or two character tokens.
	TokenType_BANG
	TokenType_BANG_EQUAL
	TokenType_EQUAL
	TokenType_EQUAL_EQUAL
	TokenType_GREATER
	TokenType_GREATER_EQUAL
	TokenType_LESS
	TokenType_LESS_EQUAL

	// Literals.
	TokenType_IDENTIFIER
	TokenType_STRING
	TokenType_NUMBER

	// Keywords.
	TokenType_AND
	TokenType_CLASS
	TokenType_ELSE
	TokenType_FALSE
	TokenType_FUN
	TokenType_FOR
	TokenType_IF
	TokenType_NIL
	TokenType_OR
	TokenType_PRINT
	TokenType_RETURN
	TokenType_SUPER
	TokenType_THIS
	TokenType_TRUE
	TokenType_VAR
	TokenType_WHILE

	TokenType_EOF
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) *Token {
	t := &Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
	return t
}
