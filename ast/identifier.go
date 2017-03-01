package ast

import (
	"fmt"

	"encoding/json"
)

// IdentifierExpression - an expression representing an identifier
type IdentifierExpression struct {
	Identifier string
}

func (e IdentifierExpression) String() string {
	return fmt.Sprintf("IDENT(%s)", e.Identifier)
}

// MarshalJSON for IdentifierExpression should just be the identifier as a string
func (e IdentifierExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Identifier)
}

