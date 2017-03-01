package ast

import (
	"fmt"
	"strconv"
	"strings"
)

// Expression - generic expression interface
type Expression interface {
	String() string
	Visit(Environment) Value
}

// Environment is the current set of all bindings and various other bits of state inside the interpreter
type Environment struct {
	Bindings map[string]Binding
}

// Binding is a name and a value
type Binding struct {
	name  string
	value *Value
}

// Value is a piece of data wrapped up with a type
type Value struct {
	Type string
	Data interface{}
}

func expressionListJoin(list []Expression, sep string) string {
	expressions := make([]string, len(list))
	for i, e := range list {
		expressions[i] = e.String()
	}
	return strings.Join(expressions, sep)
}

// NoneValue is a convenience function for returning a standardized nil value
func NoneValue() Value {
	return Value{Type: "nil", Data: nil}
}

// ErrorValue is a value for holding an error
func ErrorValue(message string) Value {
	return Value{Type: "error", Data: message}
}

func (v Value) String() string {
	switch v.Type {
	case "int":
		return strconv.FormatInt(v.Data.(int64), 10)
	case "nil":
		return "nil"
	case "error":
		return v.Data.(string)
	default:
		return fmt.Sprintf("unknown value type: %s", v.Type)
	}
}
