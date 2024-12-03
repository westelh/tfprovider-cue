package schema

import (
	"cuelang.org/go/cue/ast"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaExpr(s *tf_schema.Schema, opt Option) ast.Expr {
	switch s.Type {
	case tf_schema.TypeBool:
		return &ast.Ident{Name: "bool"}
	case tf_schema.TypeInt:
		return &ast.Ident{Name: "int"}
	case tf_schema.TypeFloat:
		return &ast.Ident{Name: "float"}
	case tf_schema.TypeString:
		return &ast.Ident{Name: "string"}

	// Expression should remark that the order of elements is preserved.
	// But cue does not regard the order of elements after all.
	case tf_schema.TypeSet: // Ordering of elements is preserved
		return listOf(s.Elem, opt)
	case tf_schema.TypeList: // Ordering of elements is NOT preserved
		return listOf(s.Elem, opt)
		
	case tf_schema.TypeMap:
		switch elem := s.Elem.(type) {
		case *tf_schema.Schema:
			return structOfSameValue(SchemaExpr(elem, opt))
		case tf_schema.ValueType:
			switch elem {
			case tf_schema.TypeBool:
				return structOfSameValue(SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeBool}, opt))
			case tf_schema.TypeInt:
				return structOfSameValue(SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeInt}, opt))
			case tf_schema.TypeFloat:
				return structOfSameValue(SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeFloat}, opt))
			case tf_schema.TypeString:
				return structOfSameValue(SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeString}, opt))
			}
		}
		return structOfSameValue(ast.NewIdent("_"))
	}
	return &ast.BadExpr{}
}

func listOf(elem interface{}, opt Option) ast.Expr {
	// TypeSet supports *Schema and *Resource elements
	switch elem := elem.(type) {
	case *tf_schema.Schema:
		return ast.NewList(&ast.Ellipsis{Type: SchemaExpr(elem, opt)})
	case *tf_schema.Resource:
		return ast.NewList(&ast.Ellipsis{Type: ResourceExpr(elem, opt)})
	default:
		return &ast.BadExpr{}
	}
}

func structOfSameValue(e ast.Expr) *ast.StructLit{
	label := ast.NewList(ast.NewIdent("_"))
	return ast.NewStruct(&ast.Field{Label: label, Value: e})
}
