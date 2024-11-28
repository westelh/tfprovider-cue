package schema

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Extracts the default value from a schema
// and returns it as a CUE ast expression
func DefaultExpr(s *tf_schema.Schema) ast.Expr {
	if s.Default == nil {
		return nil
	}

	switch s.Type {
	case tf_schema.TypeBool:
		return MarkDefault(ast.NewBool(s.Default.(bool)))
	case tf_schema.TypeInt:
		return MarkDefault(ast.NewLit(token.INT, s.Default.(string)))
	case tf_schema.TypeFloat:
		return MarkDefault(ast.NewLit(token.FLOAT, s.Default.(string)))
	case tf_schema.TypeString:
		return MarkDefault(ast.NewString(s.Default.(string)))
	}
	return nil
}
