package lexer

type StatusFunc func(lex *Lexer) StatusFunc

func detectionNumbers(data string) bool {
	base := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, v := range base {
		if data == v {
			return true
		}
	}
	return false
}

func consumptionNumbers(lex *Lexer) bool {
	checked := false
	value := ""
	for detectionNumbers(lex.peek(1)) {
		checked = true
		value += lex.Expr[lex.Index:lex.Pos]
		lex.next(1)
	}
	if checked {
		lex.appendToken(Token{
			Type:  NUMBER,
			Value: value,
		})
	}
	lex.backup()
	return checked
}

func detectionOperator(data string) bool {
	base := []string{"+", "-", "*", "/"}
	for _, v := range base {
		if data == v {
			return true
		}
	}
	return false
}

func consumpOperator(lex *Lexer) bool {
	checked := false
	value := ""
	if detectionOperator(lex.peek(1)) {
		checked = true
		value += lex.Expr[lex.Index:lex.Pos]
		lex.next(1)
	}
	if checked {
		switch value {
		case "+":
			lex.appendToken(Token{
				Type:  PLUS,
				Value: value,
			})
		case "-":
			lex.appendToken(Token{
				Type:  MINUS,
				Value: value,
			})
		case "*":
			lex.appendToken(Token{
				Type:  MUL,
				Value: value,
			})
		case "/":
			lex.appendToken(Token{
				Type:  DIV,
				Value: value,
			})
		}
	}
	lex.backup()
	return checked
}

func detectionWhite(data string) bool {
	base := []string{" "}
	for _, v := range base {
		if data == v {
			return true
		}
	}
	return false
}

func consumpWhite(lex *Lexer) bool {
	checked := false
	value := ""
	if detectionWhite(lex.peek(1)) {
		checked = true
		value += lex.Expr[lex.Index:lex.Pos]
		lex.next(1)
	}
	lex.backup()
	return checked
}

func detectionLpar(data string) bool {
	base := []string{"("}
	for _, v := range base {
		if data == v {
			return true
		}
	}
	return false
}

func consumpLpar(lex *Lexer) bool {
	checked := false
	value := ""
	if detectionLpar(lex.peek(1)) {
		checked = true
		value += lex.Expr[lex.Index:lex.Pos]
		lex.next(1)
	}
	if checked {
		lex.appendToken(Token{
			Type:  LPAR,
			Value: value,
		})
	}
	lex.backup()
	return checked
}

func detectionRpar(data string) bool {
	base := []string{")"}
	for _, v := range base {
		if data == v {
			return true
		}
	}
	return false
}

func consumpRpar(lex *Lexer) bool {
	checked := false
	value := ""
	if detectionRpar(lex.peek(1)) {
		checked = true
		value += lex.Expr[lex.Index:lex.Pos]
		lex.next(1)
	}
	if checked {
		lex.appendToken(Token{
			Type:  RPAR,
			Value: value,
		})
	}
	lex.backup()
	return checked
}
