package schema

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
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
		return MarkDefault(ast.NewLit(token.INT, strconv.Itoa(s.Default.(int))))
	case tf_schema.TypeFloat:
		// s.Default can be float32 or float64
		switch f := s.Default.(type) {
		case float32:
			return MarkDefault(ast.NewLit(token.FLOAT, strconv.FormatFloat(float64(f), 'f', -1, 32)))
		case float64:
			return MarkDefault(ast.NewLit(token.FLOAT, strconv.FormatFloat(f, 'f', -1, 64)))
		}
	case tf_schema.TypeString:
		return MarkDefault(ast.NewString(s.Default.(string)))
	}
	return nil
}
