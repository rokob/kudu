package ast

import (
	"encoding/json"
	"fmt"
)

// IntegerExpression - an expression representing an integer
type IntegerExpression struct {
	Integer int64
}

func (e IntegerExpression) String() string {
	return fmt.Sprintf("INT(%d)", e.Integer)
}

// MarshalJSON for IntegerExpression should just be the integer literal
func (e IntegerExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Integer)
}
