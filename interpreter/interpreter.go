package interpreter

import (
	"github.com/rokob/kudu/ast"
)

// Interpreter is the driver for the interpreter
type Interpreter struct {
	Environment ast.Environment
}

// New constructs an initialized Interpreter
func New() *Interpreter {
	return &Interpreter{Environment: ast.Environment{Bindings: make(map[string]ast.Binding)}}
}

// HandleExpression takes an ast expression and returns the Value based on the current environment
func (i *Interpreter) HandleExpression(expression ast.Expression) ast.Value {
	return visit(expression, i)
}

func visit(expression ast.Expression, interpreter *Interpreter) ast.Value {
	return expression.Visit(interpreter.Environment)
}
