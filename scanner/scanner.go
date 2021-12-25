package scanner

import (
	"LoxScanner/token"
	"strconv"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []*token.Token
}

func New(source string) *Scanner {
	var s = &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    0,
		tokens:  []*token.Token{},
	}
	return s
}
func (s *Scanner) ScanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.current // en cada token 'start' comienza donde 'current' termina
		s.scanToken()
	}
	// agregar EOF
	s.tokens = append(s.tokens, token.New(token.EOF, "", nil, s.line))
	return s.tokens
}

// scanToken devuelve un token
func (s *Scanner) scanToken() {
	var c byte = s.advance()
	switch c {
	case '(':
		s.addToken(token.LPAREN)
	case ')':
		s.addToken(token.RPAREN)
	case '{':
		s.addToken(token.LBRACE)
	case '}':
		s.addToken(token.RBRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.match('=') {
			s.addToken(token.NEQ)
		} else {
			s.addToken(token.NOT)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL)
		} else {
			s.addToken(token.ASSIGN)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LEQ)
		} else {
			s.addToken(token.LT)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GEQ)
		} else {
			s.addToken(token.GT)
		}
	case ' ', '\r', '\t':
		// skip white space
	case '\n':
		s.line += 1
	case '/':
		if s.match('/') {
			// skip whitespace until new line
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			// error: Unexpected character
		}
	}
}

// string avanza hasta encontrar el final del string delimitado por '"'
// considerando los saltos de líneas.
func (s *Scanner) string() {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}
	if s.isAtEnd() {
		// emitir un error aquí "Unterminated string."
		return
	}
	s.advance() // skip the closing '"'
	// extract the surrounding quotes
	var value string = s.source[s.start:s.current]
	s.addFullToken(token.STRING, value)
}

// number recorre los caracteres mientras sean de tipo number [0..9]
func (s *Scanner) number() {
	for !s.isAtEnd() && isDigit(s.peek()) {
		s.advance()
	}
	// looking for the '.'
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance() // skip the '.'
	}
	for !s.isAtEnd() && isDigit(s.peek()) {
		s.advance()
	}

	val, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic(err)
	}
	s.addFullToken(token.NUMBER, val)
}
func (s *Scanner) identifier() {
	for !s.isAtEnd() && isAlphaNumeric(s.peek()) {
		s.advance()
	}
	var text string = s.source[s.start:s.current]
	s.addToken(token.LookupIdent(text))
}

// advance incrementa la propiedad 'current' para que avance en el stream de caracteres
func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current += 1
	return c
}
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}
	s.current += 1
	return true
}

// addFullToken agrega un token solo con su tipo
func (s *Scanner) addToken(typ token.TokenType) {
	s.addFullToken(typ, nil)
}

// addToken agrega un token full
func (s *Scanner) addFullToken(typ token.TokenType, lit token.Object) {
	var text string = s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.New(typ, text, lit, s.line))
}

// peek devuelve el caracter actual sin consumirlo
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

// peekNext devuelve el siguiente caracter sin consumirlo
func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

// isAtEnd verifica si 'current' está dentro del rango de caracteres
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// isDigit verifica si el caracter dado es un número
func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// isLetter verifica si el caracter dado es una letra
func isAlpha(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

// isAlphaNumeric verifica si el caracter dado es alfanumérico
func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
