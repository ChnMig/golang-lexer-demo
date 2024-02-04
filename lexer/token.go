package lexer

// token类型
type TokenType string

const (
	NUMBER TokenType = "number" // 整数
	PLUS   TokenType = "plus"   // +
	MINUS  TokenType = "minus"  // -
	MUL    TokenType = "mul"    // *
	DIV    TokenType = "div"    // /
	LPAR   TokenType = "lpar"   // (
	RPAR   TokenType = "rpar"   // )
)

// Token 结构体
type Token struct {
	Type  TokenType
	Value string
}
