package schema

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	REQUIRED = token.NOT
	OPTIONAL = token.OPTION
)

func ResourceExpr(resource *tf_schema.Resource, dropReadOnly bool) *ast.StructLit {
	var fields map[string]*ast.Field = make(map[string]*ast.Field)
	for name, s := range resource.Schema {
		if dropReadOnly && !s.Optional && !s.Required && s.Computed {
			continue
		}
		
		f := ast.Field{}

		f.Value = SchemaExpr(s, dropReadOnly)
		f.Label = ast.NewIdent(name)

		if s.Description != "" {
			f.AddComment(singleLineComment(s.Description))
		}

		if s.Default != nil {
			f.Value = Or(MarkDefault(DefaultExpr(s)), f.Value)
		}

		if s.Optional {
			f.Constraint = OPTIONAL
		}

		if s.Required {
			f.Constraint = REQUIRED
		}

		fields[name] = &f
	}

	// ast.NewStruct(fields) takes a slice of interface{} as argument
	slice := make([]interface{}, 0, len(fields))
	for _, f := range fields {
		slice = append(slice, f)
	}

	return ast.NewStruct(slice...)
}

func singleLineComment(s string) *ast.CommentGroup {
	return &ast.CommentGroup{
		List: []*ast.Comment{{Text: s}},
	}
}
