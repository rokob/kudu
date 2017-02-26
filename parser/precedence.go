package parser

// Precedence represents the evaluation order of different operations
type Precedence int

const (
	// LOWEST represents a default for things without a well defined precedence
	LOWEST Precedence = iota
	// ASSIGNMENT - variable assignment
	ASSIGNMENT
	// CONDITIONAL - if/else
	CONDITIONAL
	// SUM - +, -
	SUM
	// PRODUCT - *, /
	PRODUCT
	// EXPONENT - ^
	EXPONENT
	// PREFIX - !, -
	PREFIX
	// POSTFIX - ++
	POSTFIX
	// CALL - function calls
	CALL
)