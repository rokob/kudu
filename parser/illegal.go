package parser

// IllegalExpression - an expression representing illegal input
type IllegalExpression struct {
	IsBreak bool
}

func (e IllegalExpression) String() string {
	return "ILLEGAL"
}
