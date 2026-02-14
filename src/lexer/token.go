package lexer

// TokenType represents the type of token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Token types
const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT  = "IDENT"  // variable names, function names
	NUMBER = "NUMBER" // 123, 45.67
	STRING = "STRING" // "hello", 'world'

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"

	// Comparison operators
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="

	// Logical operators
	BANG = "!"

	// Compound assignment
	PLUS_ASSIGN     = "+="
	MINUS_ASSIGN    = "-="
	ASTERISK_ASSIGN = "*="
	SLASH_ASSIGN    = "/="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	DOTDOTDOT = "..."

	// Keywords (Banglish)
	DHORO      = "DHORO"      // variable declaration (let/var)
	STHIR      = "STHIR"      // constant declaration (const) - স্থির = fixed/constant
	BISHWO     = "BISHWO"     // global variable declaration - বিশ্ব = world/global
	JODI       = "JODI"       // if
	NAHOLE     = "NAHOLE"     // else
	JOTOKKHON  = "JOTOKKHON"  // while
	GHURIYE    = "GHURIYE"    // for
	KAJ        = "KAJ"        // function
	FERAO      = "FERAO"      // return
	SRENI      = "SRENI"      // class (শ্রেণী)
	SHURU      = "SHURU"      // constructor (শুরু)
	NOTUN      = "NOTUN"      // new
	SOTTI      = "SOTTI"      // true
	MITTHA     = "MITTHA"     // false
	KHALI      = "KHALI"      // null
	EBONG      = "EBONG"      // and (&&)
	BA         = "BA"         // or (||)
	NA         = "NA"         // not (!)
	THAMO      = "THAMO"      // break
	CHHARO     = "CHHARO"     // continue
	ANO        = "ANO"        // import (আনো - bring)
	PATHAO     = "PATHAO"     // export (পাঠাও - send)
	HISABE     = "HISABE"     // as (হিসাবে - as/alias)
	CHESTA     = "CHESTA"     // try
	DHORO_BHUL = "DHORO_BHUL" // catch
	SHESH      = "SHESH"      // finally
	FELO       = "FELO"       // throw
	PROYASH    = "PROYASH"    // async (প্রয়াস - attempt/endeavor)
	OPEKHA     = "OPEKHA"     // await (অপেক্ষা - wait)
)

// keywords maps Banglish keywords to their token types
var keywords = map[string]TokenType{
	"dhoro":      DHORO,
	"sthir":      STHIR,
	"bishwo":     BISHWO,
	"jodi":       JODI,
	"nahole":     NAHOLE,
	"jotokkhon":  JOTOKKHON,
	"ghuriye":    GHURIYE,
	"kaj":        KAJ,
	"ferao":      FERAO,
	"sreni":      SRENI,
	"shuru":      SHURU,
	"notun":      NOTUN,
	"sotti":      SOTTI,
	"mittha":     MITTHA,
	"khali":      KHALI,
	"ebong":      EBONG,
	"ba":         BA,
	"na":         NA,
	"thamo":      THAMO,
	"chharo":     CHHARO,
	"ano":        ANO,
	"pathao":     PATHAO,
	"hisabe":     HISABE,
	"chesta":     CHESTA,
	"dhoro_bhul": DHORO_BHUL,
	"shesh":      SHESH,
	"felo":       FELO,
	"proyash":    PROYASH,
	"opekha":     OPEKHA,
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// NewToken creates a new token
func NewToken(tokenType TokenType, literal string, line, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
	}
}
