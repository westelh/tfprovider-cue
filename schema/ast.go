package schema

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
)

func MarkDefault(x ast.Expr) ast.Expr {
	return &ast.UnaryExpr{
		Op: token.MUL,
		X:  x,
	}
}

func Or(x ast.Expr, y ast.Expr) ast.Expr {
	return &ast.BinaryExpr{
		X:  x,
		Op: token.OR,
		Y:  y,
	}
}
