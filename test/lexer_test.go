package test

import (
	"BanglaCode/src/lexer"
	"testing"
)

func TestNextToken_Keywords(t *testing.T) {
	input := `dhoro jodi nahole jotokkhon ghuriye kaj ferao sreni shuru notun sotti mittha khali ebong ba na thamo chharo ano pathao hisabe chesta dhoro_bhul shesh felo`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.DHORO, "dhoro"},
		{lexer.JODI, "jodi"},
		{lexer.NAHOLE, "nahole"},
		{lexer.JOTOKKHON, "jotokkhon"},
		{lexer.GHURIYE, "ghuriye"},
		{lexer.KAJ, "kaj"},
		{lexer.FERAO, "ferao"},
		{lexer.SRENI, "sreni"},
		{lexer.SHURU, "shuru"},
		{lexer.NOTUN, "notun"},
		{lexer.SOTTI, "sotti"},
		{lexer.MITTHA, "mittha"},
		{lexer.KHALI, "khali"},
		{lexer.EBONG, "ebong"},
		{lexer.BA, "ba"},
		{lexer.NA, "na"},
		{lexer.THAMO, "thamo"},
		{lexer.CHHARO, "chharo"},
		{lexer.ANO, "ano"},
		{lexer.PATHAO, "pathao"},
		{lexer.HISABE, "hisabe"},
		{lexer.CHESTA, "chesta"},
		{lexer.DHORO_BHUL, "dhoro_bhul"},
		{lexer.SHESH, "shesh"},
		{lexer.FELO, "felo"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Operators(t *testing.T) {
	input := `= + - * / % == != < > <= >= ! += -= *= /=`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.ASSIGN, "="},
		{lexer.PLUS, "+"},
		{lexer.MINUS, "-"},
		{lexer.ASTERISK, "*"},
		{lexer.SLASH, "/"},
		{lexer.PERCENT, "%"},
		{lexer.EQ, "=="},
		{lexer.NOT_EQ, "!="},
		{lexer.LT, "<"},
		{lexer.GT, ">"},
		{lexer.LTE, "<="},
		{lexer.GTE, ">="},
		{lexer.BANG, "!"},
		{lexer.PLUS_ASSIGN, "+="},
		{lexer.MINUS_ASSIGN, "-="},
		{lexer.ASTERISK_ASSIGN, "*="},
		{lexer.SLASH_ASSIGN, "/="},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Delimiters(t *testing.T) {
	input := `, ; : . ( ) { } [ ]`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.COMMA, ","},
		{lexer.SEMICOLON, ";"},
		{lexer.COLON, ":"},
		{lexer.DOT, "."},
		{lexer.LPAREN, "("},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.RBRACE, "}"},
		{lexer.LBRACKET, "["},
		{lexer.RBRACKET, "]"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Numbers(t *testing.T) {
	input := `123 45.67 0 3.14159 100`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.NUMBER, "123"},
		{lexer.NUMBER, "45.67"},
		{lexer.NUMBER, "0"},
		{lexer.NUMBER, "3.14159"},
		{lexer.NUMBER, "100"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Strings(t *testing.T) {
	input := `"hello" 'world' "hello world" 'with spaces'`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.STRING, "hello"},
		{lexer.STRING, "world"},
		{lexer.STRING, "hello world"},
		{lexer.STRING, "with spaces"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Identifiers(t *testing.T) {
	input := `foo bar myVar some_variable _underscore`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.IDENT, "foo"},
		{lexer.IDENT, "bar"},
		{lexer.IDENT, "myVar"},
		{lexer.IDENT, "some_variable"},
		{lexer.IDENT, "_underscore"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_VariableDeclaration(t *testing.T) {
	input := `dhoro x = 5;`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.DHORO, "dhoro"},
		{lexer.IDENT, "x"},
		{lexer.ASSIGN, "="},
		{lexer.NUMBER, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_IfElseStatement(t *testing.T) {
	input := `jodi (x > 5) { ferao sotti; } nahole { ferao mittha; }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.JODI, "jodi"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "x"},
		{lexer.GT, ">"},
		{lexer.NUMBER, "5"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.FERAO, "ferao"},
		{lexer.SOTTI, "sotti"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.NAHOLE, "nahole"},
		{lexer.LBRACE, "{"},
		{lexer.FERAO, "ferao"},
		{lexer.MITTHA, "mittha"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_FunctionDeclaration(t *testing.T) {
	input := `kaj add(a, b) { ferao a + b; }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.KAJ, "kaj"},
		{lexer.IDENT, "add"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "a"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "b"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.FERAO, "ferao"},
		{lexer.IDENT, "a"},
		{lexer.PLUS, "+"},
		{lexer.IDENT, "b"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_WhileLoop(t *testing.T) {
	input := `jotokkhon (x < 10) { x = x + 1; }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.JOTOKKHON, "jotokkhon"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "x"},
		{lexer.LT, "<"},
		{lexer.NUMBER, "10"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.IDENT, "x"},
		{lexer.ASSIGN, "="},
		{lexer.IDENT, "x"},
		{lexer.PLUS, "+"},
		{lexer.NUMBER, "1"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_ForLoop(t *testing.T) {
	input := `ghuriye (dhoro i = 0; i < 5; i = i + 1) { }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.GHURIYE, "ghuriye"},
		{lexer.LPAREN, "("},
		{lexer.DHORO, "dhoro"},
		{lexer.IDENT, "i"},
		{lexer.ASSIGN, "="},
		{lexer.NUMBER, "0"},
		{lexer.SEMICOLON, ";"},
		{lexer.IDENT, "i"},
		{lexer.LT, "<"},
		{lexer.NUMBER, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.IDENT, "i"},
		{lexer.ASSIGN, "="},
		{lexer.IDENT, "i"},
		{lexer.PLUS, "+"},
		{lexer.NUMBER, "1"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_ClassDeclaration(t *testing.T) {
	input := `sreni Person { shuru(naam) { ei.naam = naam; } }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.SRENI, "sreni"},
		{lexer.IDENT, "Person"},
		{lexer.LBRACE, "{"},
		{lexer.SHURU, "shuru"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "naam"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.IDENT, "ei"},
		{lexer.DOT, "."},
		{lexer.IDENT, "naam"},
		{lexer.ASSIGN, "="},
		{lexer.IDENT, "naam"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_TryCatch(t *testing.T) {
	input := `chesta { felo "error"; } dhoro_bhul (e) { } shesh { }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.CHESTA, "chesta"},
		{lexer.LBRACE, "{"},
		{lexer.FELO, "felo"},
		{lexer.STRING, "error"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.DHORO_BHUL, "dhoro_bhul"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "e"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.RBRACE, "}"},
		{lexer.SHESH, "shesh"},
		{lexer.LBRACE, "{"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_ArrayAndMap(t *testing.T) {
	input := `[1, 2, 3] {naam: "Ankan", boyes: 25}`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.LBRACKET, "["},
		{lexer.NUMBER, "1"},
		{lexer.COMMA, ","},
		{lexer.NUMBER, "2"},
		{lexer.COMMA, ","},
		{lexer.NUMBER, "3"},
		{lexer.RBRACKET, "]"},
		{lexer.LBRACE, "{"},
		{lexer.IDENT, "naam"},
		{lexer.COLON, ":"},
		{lexer.STRING, "Ankan"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "boyes"},
		{lexer.COLON, ":"},
		{lexer.NUMBER, "25"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Comments(t *testing.T) {
	input := `dhoro x = 5; // this is a comment
dhoro y = 10;`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.DHORO, "dhoro"},
		{lexer.IDENT, "x"},
		{lexer.ASSIGN, "="},
		{lexer.NUMBER, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.DHORO, "dhoro"},
		{lexer.IDENT, "y"},
		{lexer.ASSIGN, "="},
		{lexer.NUMBER, "10"},
		{lexer.SEMICOLON, ";"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_LogicalOperators(t *testing.T) {
	input := `sotti ebong mittha ba na sotti`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.SOTTI, "sotti"},
		{lexer.EBONG, "ebong"},
		{lexer.MITTHA, "mittha"},
		{lexer.BA, "ba"},
		{lexer.NA, "na"},
		{lexer.SOTTI, "sotti"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_BreakContinue(t *testing.T) {
	input := `thamo; chharo;`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.THAMO, "thamo"},
		{lexer.SEMICOLON, ";"},
		{lexer.CHHARO, "chharo"},
		{lexer.SEMICOLON, ";"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_ImportExport(t *testing.T) {
	input := `ano "math.bang" hisabe math; pathao kaj add(a, b) { ferao a + b; }`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.ANO, "ano"},
		{lexer.STRING, "math.bang"},
		{lexer.HISABE, "hisabe"},
		{lexer.IDENT, "math"},
		{lexer.SEMICOLON, ";"},
		{lexer.PATHAO, "pathao"},
		{lexer.KAJ, "kaj"},
		{lexer.IDENT, "add"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "a"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "b"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.FERAO, "ferao"},
		{lexer.IDENT, "a"},
		{lexer.PLUS, "+"},
		{lexer.IDENT, "b"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_NewExpression(t *testing.T) {
	input := `notun Person("Ankan")`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.NOTUN, "notun"},
		{lexer.IDENT, "Person"},
		{lexer.LPAREN, "("},
		{lexer.STRING, "Ankan"},
		{lexer.RPAREN, ")"},
		{lexer.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_LineAndColumn(t *testing.T) {
	input := `dhoro x = 5;
dhoro y = 10;`

	l := lexer.New(input)

	// First line - dhoro
	tok := l.NextToken()
	if tok.Line != 1 || tok.Column != 1 {
		t.Errorf("expected line=1, column=1, got line=%d, column=%d", tok.Line, tok.Column)
	}

	// x
	tok = l.NextToken()
	if tok.Line != 1 || tok.Column != 7 {
		t.Errorf("expected line=1, column=7, got line=%d, column=%d", tok.Line, tok.Column)
	}

	// Skip to second line
	l.NextToken() // =
	l.NextToken() // 5
	l.NextToken() // ;
	tok = l.NextToken() // dhoro on line 2
	if tok.Line != 2 {
		t.Errorf("expected line=2, got line=%d", tok.Line)
	}
}

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		ident    string
		expected lexer.TokenType
	}{
		{"dhoro", lexer.DHORO},
		{"jodi", lexer.JODI},
		{"nahole", lexer.NAHOLE},
		{"kaj", lexer.KAJ},
		{"ferao", lexer.FERAO},
		{"sotti", lexer.SOTTI},
		{"mittha", lexer.MITTHA},
		{"khali", lexer.KHALI},
		{"myVariable", lexer.IDENT},
		{"foo", lexer.IDENT},
	}

	for _, tt := range tests {
		got := lexer.LookupIdent(tt.ident)
		if got != tt.expected {
			t.Errorf("LookupIdent(%q) = %q, want %q", tt.ident, got, tt.expected)
		}
	}
}
