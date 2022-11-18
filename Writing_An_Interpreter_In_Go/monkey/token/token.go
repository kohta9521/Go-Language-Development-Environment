// token/token.go

package token

type TokenType string

type Tokn struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//  識別子 + リテラル
	INDENT = "INDENT" // add, foovar. x. y
	INT    = "INT"    // 1343456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLOM = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
