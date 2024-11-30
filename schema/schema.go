package schema

import (
	"cuelang.org/go/cue/ast"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaExpr(s *tf_schema.Schema, dropReadOnly bool) ast.Expr {
	switch s.Type {
	case tf_schema.TypeBool:
		return &ast.Ident{Name: "bool"}
	case tf_schema.TypeInt:
		return &ast.Ident{Name: "int"}
	case tf_schema.TypeFloat:
		return &ast.Ident{Name: "float"}
	case tf_schema.TypeString:
		return &ast.Ident{Name: "string"}
	case tf_schema.TypeSet: // Ordering of elements is preserved
		// Expression should remark that the order of elements is preserved.
		// But cue does not regard the order of elements after all.
		return listOf(s.Elem, dropReadOnly)
	case tf_schema.TypeList: // Ordering of elements is NOT preserved
		return listOf(s.Elem, dropReadOnly)
	case tf_schema.TypeMap:
		switch elem := s.Elem.(type) {
		case *tf_schema.Schema:
			return mapOf(elem, dropReadOnly)
		case tf_schema.ValueType:
			switch elem {
			case tf_schema.TypeBool:
				return mapOf(&tf_schema.Schema{Type: tf_schema.TypeBool}, dropReadOnly)
			case tf_schema.TypeInt:
				return mapOf(&tf_schema.Schema{Type: tf_schema.TypeInt}, dropReadOnly)
			case tf_schema.TypeFloat:
				return mapOf(&tf_schema.Schema{Type: tf_schema.TypeFloat}, dropReadOnly)
			case tf_schema.TypeString:
				return mapOf(&tf_schema.Schema{Type: tf_schema.TypeString}, dropReadOnly)
			}
			return nil
		}
	}
	return nil
}

func listOf(elem interface{}, dropReadOnly bool) ast.Expr {
	// TypeSet supports *Schema and *Resource elements
	switch elem := elem.(type) {
	case *tf_schema.Schema:
		return ast.NewList(&ast.Ellipsis{Type: SchemaExpr(elem, dropReadOnly)})
	case *tf_schema.Resource:
		return ast.NewList(&ast.Ellipsis{Type: ResourceExpr(elem, dropReadOnly)})
	}
	return nil
}

func mapOf(s *tf_schema.Schema, dropReadOnly bool) ast.Expr {
	label := ast.NewList(ast.NewIdent("_"))
	return ast.NewStruct(&ast.Field{Label: label, Value: SchemaExpr(s, dropReadOnly)})
}
