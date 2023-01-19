package lox

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

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {

}
