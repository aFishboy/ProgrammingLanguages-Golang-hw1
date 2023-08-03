package simplify

import (
	"fmt"
	"hw1/expr"
)

// Simplify should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	// Use a switch statement to handle different types of expressions
	switch e := e.(type) {
	// If the expression is a Literal, return it as is
	case expr.Literal:
		return e
	// If the expression is a Var, check if it's in the environment and return its value if it is
	case expr.Var:
		if val, notEmpty := env[e]; notEmpty {
			return expr.Literal(val)
		}
		return e
	// If the expression is a Unary, simplify its child expression and return the simplified result
	case expr.Unary:
		exp := Simplify(e.X, env)
		if literalVal, isLiteral := exp.(expr.Literal); isLiteral {
			// Handle unary operations on literals
			switch e.Op {
			case '+':
				return literalVal
			case '-':
				return expr.Literal(-literalVal)
			}
		}
		return expr.Unary{Op: e.Op, X: exp}
	// If the expression is a Binary, simplify its left and right child expressions and return the simplified result
	case expr.Binary:
		left := Simplify(e.X, env)
		right := Simplify(e.Y, env)
		// Check for cases where one or both operands are literals
		leftLit, leftIsLit := left.(expr.Literal)
		rightLit, rightIsLit := right.(expr.Literal)
		if leftIsLit && rightIsLit {
			// Handle binary operations where both operands are Literals
			switch e.Op {
			case '+':
				return expr.Literal(leftLit + rightLit)
			case '-':
				return expr.Literal(leftLit - rightLit)
			case '*':
				return expr.Literal(leftLit * rightLit)
			case '/':
				return expr.Literal(leftLit / rightLit)
			}
		} else if leftIsLit && leftLit == 0 {
			// Handle binary operations with a 0 operand
			switch e.Op {
			case '+':
				return right
			case '*':
				return expr.Literal(0)
			}
		} else if rightIsLit && rightLit == 0 {
			// Handle binary operations with a 0 operand
			switch e.Op {
			case '+':
				return left
			case '*':
				return expr.Literal(0)
			}
		} else if leftIsLit && leftLit == 1 {
			// Handle binary operations with a 1 operand
			switch e.Op {
			case '*':
				return right
			}
		} else if rightIsLit && rightLit == 1 {
			// Handle binary operations with a 1 operand
			switch e.Op {
			case '*':
				return left
			}
		} else {
			// Return the original binary expression
			return expr.Binary{X: left, Op: e.Op, Y: right}
		}
	}
	// If the expression is of an unknown type, panic and print an error message
	panic(fmt.Sprintf("unknown Expr: %T", e))
}
