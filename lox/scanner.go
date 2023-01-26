package lox

import (
	"fmt"
	"strconv"
)

var keywords = map[string]TokenType{
	"and":    TokenType_AND,
	"class":  TokenType_CLASS,
	"else":   TokenType_ELSE,
	"false":  TokenType_FALSE,
	"for":    TokenType_FOR,
	"fun":    TokenType_FUN,
	"if":     TokenType_IF,
	"nil":    TokenType_NIL,
	"or":     TokenType_OR,
	"print":  TokenType_PRINT,
	"return": TokenType_RETURN,
	"super":  TokenType_SUPER,
	"this":   TokenType_THIS,
	"true":   TokenType_TRUE,
	"var":    TokenType_VAR,
	"while":  TokenType_WHILE,
}

type Scanner struct {
	source  string
	tokens  []*Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	s := &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
	return s
}

func (s *Scanner) scanTokens() []*Token {

	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(TokenType_EOF, "", nil, s.line))

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(TokenType_LEFT_PAREN, nil)
	case ')':
		s.addToken(TokenType_RIGHT_PAREN, nil)
	case '{':
		s.addToken(TokenType_LEFT_BRACE, nil)
	case '}':
		s.addToken(TokenType_RIGHT_BRACE, nil)
	case ',':
		s.addToken(TokenType_COMMA, nil)
	case '.':
		s.addToken(TokenType_DOT, nil)
	case '-':
		s.addToken(TokenType_MINUS, nil)
	case '+':
		s.addToken(TokenType_PLUS, nil)
	case ';':
		s.addToken(TokenType_SEMICOLON, nil)
	case '*':
		s.addToken(TokenType_STAR, nil)
	case '!':
		var tokenType TokenType
		if s.match('=') {
			tokenType = TokenType_BANG_EQUAL
		} else {
			tokenType = TokenType_BANG
		}
		s.addToken(tokenType, nil)
	case '=':
		var tokenType TokenType
		if s.match('=') {
			tokenType = TokenType_EQUAL_EQUAL
		} else {
			tokenType = TokenType_EQUAL
		}
		s.addToken(tokenType, nil)
	case '<':
		var tokenType TokenType
		if s.match('=') {
			tokenType = TokenType_LESS_EQUAL
		} else {
			tokenType = TokenType_LESS
		}
		s.addToken(tokenType, nil)
	case '>':
		var tokenType TokenType
		if s.match('=') {
			tokenType = TokenType_GREATER_EQUAL
		} else {
			tokenType = TokenType_GREATER
		}
		s.addToken(tokenType, nil)
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(TokenType_SLASH, nil)
		}
	case '\n':
		s.line++
	case ' ', '\r', '\t':
	// Ignore whitespace.
	case '"':
		err := s.string()
		if err != nil {
			panic(err)
		}
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			panic(fmt.Sprintf("[%d]Unexpected character.", s.line))
		}
	}
}

func (s *Scanner) advance() uint8 {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) peek() uint8 {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() uint8 {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) match(expected uint8) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		return fmt.Errorf("[%d]Unterminated string", s.line)
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]
	s.addToken(TokenType_STRING, value)
	return nil
}

func (s *Scanner) isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c uint8) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *Scanner) isAlphaNumeric(c uint8) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part.
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}

	}

	text := s.source[s.start:s.current]

	value, _ := strconv.ParseFloat(text, 64)

	s.addToken(TokenType_NUMBER, value)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = TokenType_IDENTIFIER
	}
	s.addToken(tokenType, nil)
}
