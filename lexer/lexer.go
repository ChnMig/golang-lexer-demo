package lexer

import "container/list"

// Lexer 结构体
type Lexer struct {
	Expr   string     // 源表达式
	Tokens []Token    // 生成的Token
	Pos    int        // 当前读取到的位置
	Index  int        // 真正处理的索引所在位置
	List   *list.List // 双向链表,存储AST中间信息
	Node   *AstNode   // AST树
}

func (l *Lexer) peek(length int) string {
	if l.Pos+length > len(l.Expr) {
		return ""
	}
	data := string(l.Expr[l.Pos : l.Pos+length])
	l.Pos = l.Pos + length
	return data
}

func (l *Lexer) next(length int) string {
	if l.Index+length > len(l.Expr) {
		return ""
	}
	data := string(l.Expr[l.Index : l.Index+length])
	l.Index = l.Index + length
	return data
}

func (l *Lexer) backup() {
	l.Pos = l.Index
}

func (l *Lexer) appendToken(t Token) {
	l.Tokens = append(l.Tokens, t)
}
