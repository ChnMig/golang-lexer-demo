package lexer

// Node 结构体
type AstNode struct {
	Type  TokenType
	Value string

	Left  *AstNode
	Right *AstNode
}

func NewAstNode(typ TokenType, value string) *AstNode {
	return &AstNode{
		Type:  typ,
		Value: value,
	}
}
