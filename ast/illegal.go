package ast

// IllegalExpression - an expression representing illegal input
type IllegalExpression struct {
	IsBreak bool
}

func (e IllegalExpression) String() string {
	return "ILLEGAL"
}

// Visit this AST Node
func (e IllegalExpression) Visit(env Environment) Value {
	return NoneValue()
}
