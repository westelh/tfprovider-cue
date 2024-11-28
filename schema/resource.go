package schema

import (
	"cuelang.org/go/cue/ast"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceExpr(resource *tf_schema.Resource) *ast.StructLit {
	var fields map[string]*ast.Field = make(map[string]*ast.Field)
	for name, s := range resource.Schema {
		f := ast.Field{}

		f.Value = SchemaExpr(s)
		f.Label = ast.NewIdent(name)

		if s.Description != "" {
			f.AddComment(singleLineComment(s.Description))
		}

		if s.Default != nil {
			f.Value = Or(MarkDefault(DefaultExpr(s)), f.Value)
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