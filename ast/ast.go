package ast

import (
	"strings"
)

// Expression - generic expression interface
type Expression interface {
	String() string
}

func expressionListJoin(list []Expression, sep string) string {
	expressions := make([]string, len(list))
	for i, e := range list {
		expressions[i] = e.String()
	}
	return strings.Join(expressions, sep)
}
