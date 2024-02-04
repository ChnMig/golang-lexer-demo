package lexer

import (
	"container/list"
	"fmt"
)

func precedence(typ TokenType) int {
	switch typ {
	case PLUS:
		fallthrough
	case MINUS:
		return 1
	case MUL:
		fallthrough
	case DIV:
		return 2
	default:
		return -1
	}
}

func (l *Lexer) CreateList() error {
	opStack := NewStack() // 存放()内数据
	postFix := list.New() // 存放处理后的节点信息
	for _, t := range l.Tokens {
		// 是数字直接放入postFix
		if t.Type == NUMBER {
			postFix.PushBack(t)
			continue
		}
		// 是操作符( 直接放入opStack, 代表开始处理()内数据
		if t.Type == LPAR {
			opStack.Push(t)
			continue
		}
		// 是操作符) 代表结束处理()内数据
		if t.Type == RPAR {
			// 从栈内弹出数据到链表
			for opStack.Len() > 0 && opStack.Top().(Token).Type != LPAR {
				postFix.PushBack(opStack.Pop())
			}
			// 处理表达式异常的情况
			if opStack.Len() > 0 && opStack.Top().(Token).Type != LPAR {
				return fmt.Errorf("mismatched parenthesis")
			}
			if opStack.Len() == 0 {
				return fmt.Errorf("mismatched parenthesis")
			}
			opStack.Pop()
		} else {
			// 是操作符, 优先级比栈顶元素小, 则弹出栈顶元素
			for opStack.Len() > 0 && precedence(t.Type) <= precedence(opStack.Top().(Token).Type) {
				postFix.PushBack(opStack.Pop())
			}
			opStack.Push(t)
		}
	}
	for opStack.Len() > 0 {
		postFix.PushBack(opStack.Pop())
	}
	l.List = postFix
	return nil
}

func (l *Lexer) ConstructAst() error {
	stack := NewStack()
	for item := l.List.Front(); item != nil; item = item.Next() {
		lexItem := item.Value.(Token)
		if lexItem.Type == NUMBER {
			stack.Push(NewAstNode(NUMBER, lexItem.Value))
		} else {
			node := NewAstNode(lexItem.Type, lexItem.Value)
			if stack.Len() < 2 {
				return fmt.Errorf("invalid expression")
			}
			// 先进后出, 防止结合性问题
			node.Right = stack.Pop().(*AstNode)
			node.Left = stack.Pop().(*AstNode)
			stack.Push(node)
		}
	}
	if stack.Len() < 1 {
		return fmt.Errorf("invalid expression")
	}
	l.Node = stack.Pop().(*AstNode)
	return nil
}

func (l *Lexer) Interpreter() (int, error) {
	return interpreterNode(l.Node)
}
