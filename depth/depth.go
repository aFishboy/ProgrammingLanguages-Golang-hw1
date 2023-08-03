package depth

import (
	"hw1/expr"
	"fmt"
)

// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	/* switch to determine type of input*/
	switch e := e.(type) {
	case expr.Literal, expr.Var: // if expression is literal or variable
		return 1
	case expr.Unary: // expression is Unaray
		return Depth(e.X) + 1 // call depth on the child expression and add that to 1
	case expr.Binary:
		lDepth := Depth(e.X) // call depth for left child expression
		rDepth := Depth(e.Y) // call depth for right child expression
		if lDepth > rDepth { // determine which has greater depth
			return lDepth + 1
		} else{
			return rDepth + 1
		}
	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}

