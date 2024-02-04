package lexer

func exprParse(lex *Lexer) StatusFunc {
	switch {
	case consumptionNumbers(lex):
		return exprParse
	case consumpOperator(lex):
		return exprParse
	case consumpWhite(lex):
		return exprParse
	case consumpLpar(lex):
		return exprParse
	case consumpRpar(lex):
		return exprParse
	default:
		return nil
	}
}

func Parse(expr string) (Lexer, error) {
	l := &Lexer{
		Expr: expr,
	}
	for p := exprParse; p != nil; {
		p = p(l)
	}
	return *l, nil
}
