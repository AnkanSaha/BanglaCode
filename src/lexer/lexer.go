package lexer

import (
	"unicode"
)

// Lexer represents the lexical analyzer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar advances the lexer position and updates current character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" (end of input)
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++

	// Track newlines for error reporting
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar returns the next character without advancing position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	// Handle comments
	if l.ch == '/' && l.peekChar() == '/' {
		l.skipComment()
		return l.NextToken()
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(EQ, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(ASSIGN, string(l.ch), l.line, l.column)
		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(PLUS_ASSIGN, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(PLUS, string(l.ch), l.line, l.column)
		}
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(MINUS_ASSIGN, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(MINUS, string(l.ch), l.line, l.column)
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(ASTERISK_ASSIGN, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(ASTERISK, string(l.ch), l.line, l.column)
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(SLASH_ASSIGN, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(SLASH, string(l.ch), l.line, l.column)
		}
	case '%':
		tok = NewToken(PERCENT, string(l.ch), l.line, l.column)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(NOT_EQ, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(BANG, string(l.ch), l.line, l.column)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(LTE, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(LT, string(l.ch), l.line, l.column)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			line := l.line
			column := l.column
			l.readChar()
			tok = NewToken(GTE, string(ch)+string(l.ch), line, column)
		} else {
			tok = NewToken(GT, string(l.ch), l.line, l.column)
		}
	case ',':
		tok = NewToken(COMMA, string(l.ch), l.line, l.column)
	case ';':
		tok = NewToken(SEMICOLON, string(l.ch), l.line, l.column)
	case ':':
		tok = NewToken(COLON, string(l.ch), l.line, l.column)
	case '.':
		if l.peekChar() == '.' && l.readPosition+1 < len(l.input) && l.input[l.readPosition+1] == '.' {
			line := l.line
			column := l.column
			l.readChar() // consume second '.'
			l.readChar() // consume third '.'
			tok = NewToken(DOTDOTDOT, "...", line, column)
		} else {
			tok = NewToken(DOT, string(l.ch), l.line, l.column)
		}
	case '(':
		tok = NewToken(LPAREN, string(l.ch), l.line, l.column)
	case ')':
		tok = NewToken(RPAREN, string(l.ch), l.line, l.column)
	case '{':
		tok = NewToken(LBRACE, string(l.ch), l.line, l.column)
	case '}':
		tok = NewToken(RBRACE, string(l.ch), l.line, l.column)
	case '[':
		tok = NewToken(LBRACKET, string(l.ch), l.line, l.column)
	case ']':
		tok = NewToken(RBRACKET, string(l.ch), l.line, l.column)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString('"')
		tok.Line = l.line
		tok.Column = l.column
		l.readChar() // advance past closing quote
		return tok
	case '\'':
		tok.Type = STRING
		tok.Literal = l.readString('\'')
		tok.Line = l.line
		tok.Column = l.column
		l.readChar() // advance past closing quote
		return tok
	case '`':
		tok.Type = TEMPLATE
		tok.Literal = l.readTemplate()
		tok.Line = l.line
		tok.Column = l.column
		l.readChar() // advance past closing backtick
		return tok
	case 0:
		tok = NewToken(EOF, "", l.line, l.column)
	default:
		if isLetter(l.ch) {
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal = l.readNumber()
			tok.Type = NUMBER
			return tok
		} else {
			tok = NewToken(ILLEGAL, string(l.ch), l.line, l.column)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads an identifier (variable name, keyword, etc.)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a numeric literal (integer or float)
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	// Handle decimal numbers
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar() // consume '.'
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

// readString reads a string literal
func (l *Lexer) readString(quote byte) string {
	position := l.position + 1 // skip opening quote
	for {
		l.readChar()
		if l.ch == quote || l.ch == 0 {
			break
		}
		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // skip escaped character
		}
	}
	str := l.input[position:l.position]
	return str
}

// readTemplate reads a template literal with ${expression} interpolation
func (l *Lexer) readTemplate() string {
	position := l.position + 1 // skip opening backtick
	braceDepth := 0

	for {
		l.readChar()
		if l.ch == 0 {
			break // end of input
		}

		// Track brace depth for ${...} expressions
		if l.ch == '$' && l.peekChar() == '{' {
			l.readChar() // consume '{'
			braceDepth++
		} else if l.ch == '{' && braceDepth > 0 {
			braceDepth++
		} else if l.ch == '}' && braceDepth > 0 {
			braceDepth--
		} else if l.ch == '`' && braceDepth == 0 {
			break // found closing backtick (not in expression)
		}

		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // skip escaped character
		}
	}

	str := l.input[position:l.position]
	return str
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips single-line comments
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// isLetter checks if a character is a letter or underscore
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// isDigit checks if a character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
