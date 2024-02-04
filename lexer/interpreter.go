package lexer

import (
	"fmt"
	"strconv"
)

// 真正的计算函数
type arithmeticFunc func(int, int) int

var interpretMap map[TokenType]arithmeticFunc

func plus(left, right int) int {
	return left + right
}
func minus(left, right int) int {
	return left - right
}
func mul(left, right int) int {
	return left * right
}
func div(left, right int) int {
	return left / right
}

func interpreterNode(node *AstNode) (int, error) {
	if node.Value == "" {
		return 0, fmt.Errorf("expected value, got nil")
	}
	// 没有子节点了
	if node.Left == nil && node.Right == nil {
		if node.Value == "" {
			return 0, fmt.Errorf("expected value, got nil")
		}
		number, err := strconv.Atoi(node.Value)
		if err != nil {
			return 0, fmt.Errorf("expected number, got %s", node.Value)
		}
		return number, nil
	}
	// 有子节点, 递归探测
	fun, ok := interpretMap[node.Type]
	if !ok {
		return 0, fmt.Errorf("unknown type %s", node.Type)
	}
	// 先左后右
	left, err := interpreterNode(node.Left)
	if err != nil {
		return 0, err
	}
	right, err := interpreterNode(node.Right)
	if err != nil {
		return 0, err
	}
	// 计算
	return fun(left, right), nil
}

func init() {
	interpretMap = make(map[TokenType]arithmeticFunc)
	interpretMap[PLUS] = plus
	interpretMap[MINUS] = minus
	interpretMap[MUL] = mul
	interpretMap[DIV] = div
}
