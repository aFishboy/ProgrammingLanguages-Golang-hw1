package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	"strconv"
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	ast.Inspect(node, func(n ast.Node) bool {
		/* type assertion of the node to check if it is type *ast.CallExpr */
		callExpr, valid := n.(*ast.CallExpr)
		if !valid {
			return true
		}
		/* check if it is type *ast.SelectorExpr */
		funcIdent, valid := callExpr.Fun.(*ast.SelectorExpr)
		if !valid {
			return true
		}

		/* check if it is named ParseAndEval */
		if funcIdent.Sel.Name != "ParseAndEval" {
			return true
		}

		if len(callExpr.Args) != 2 {
			return true
		}

		/* Check if the first argument is a string literal */
		lit, valid := callExpr.Args[0].(*ast.BasicLit)
		if !valid || lit.Kind != token.STRING {
			return true
		}

		/* strconv.Unquote to turn the ast.Basic to string and remove quotes */
		exprString, _ := strconv.Unquote(lit.Value)

		/* use expr.Parse to turn string into expr.Expr var */
		parsedExpr, err1 := expr.Parse(exprString)
		if err1 != nil {
			return true
		}

		/* call simplify on parsedExpr with empty env */
		simplifiedExpr := simplify.Simplify(parsedExpr, expr.Env{})
		/* turn into a string and add quotes back to string */
		simplifiedStr := "\"" + expr.Format(simplifiedExpr) + "\""

		/* create an ast.BasicLit with string value of simplifiedStr*/
		simplifiedAst := &ast.BasicLit{
			Kind:  token.INT,
			Value: simplifiedStr,
		}

		/* put BasicLit back into ast */
		callExpr.Args[0] = simplifiedAst
		return true
	})
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic("was error for parseFile")
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
