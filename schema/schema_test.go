package schema_test

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/westelh/tfprovider-cue/schema"
	"reflect"
	"testing"
)

var formatOpt format.Option = format.Simplify()

func FormatNode(n *ast.Node) []byte {
	b, _ := format.Node(*n, formatOpt)
	return b
}

func FormatExpr(e *ast.Expr) []byte {
	b, _ := format.Node(*e, formatOpt)
	return b
}

func FormatString(s string) []byte {
	b, _ := format.Source([]byte(s), formatOpt)
	return b
}

func TestBoolExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeBool}, schema.Option{DropReadOnly: false})
	want := "bool"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestIntExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeInt}, schema.Option{DropReadOnly: false})
	want := "int"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestFloatExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeFloat}, schema.Option{DropReadOnly: false})
	want := "float"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestStringExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{Type: tf_schema.TypeString}, schema.Option{DropReadOnly: false})
	want := "string"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfSetExpr(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"set": {Type: tf_schema.TypeSet, Elem: &tf_schema.Schema{Type: tf_schema.TypeString}},
		},
	}, schema.Option{DropReadOnly: false})
	want := "{set:[...string]}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfListExpr(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"list": {Type: tf_schema.TypeList, Elem: &tf_schema.Schema{Type: tf_schema.TypeString}},
		},
	}, schema.Option{DropReadOnly: false})
	want := "{list:[...string]}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestSetOfMapExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeSet,
		Elem: &tf_schema.Resource{
			Schema: map[string]*tf_schema.Schema{
				"foo": {Type: tf_schema.TypeString},
				"bar": {Type: tf_schema.TypeInt},
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := "[{foo: string, bar: int}]"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestSetOfListExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeSet,
		Elem: &tf_schema.Schema{Type: tf_schema.TypeList, Elem: &tf_schema.Schema{Type: tf_schema.TypeString}},
	}, schema.Option{DropReadOnly: false})
	want := "[[...string]]"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestListOfMapExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeList,
		Elem: &tf_schema.Resource{
			Schema: map[string]*tf_schema.Schema{
				"foo": {Type: tf_schema.TypeString},
				"bar": {Type: tf_schema.TypeInt},
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := "[...{foo:string, bar:int}]"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestListOfSetExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeList,
		Elem: &tf_schema.Schema{Type: tf_schema.TypeSet, Elem: &tf_schema.Schema{Type: tf_schema.TypeString}},
	}, schema.Option{DropReadOnly: false})
	want := "[[...string]]"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfValueTypeBoolExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeMap,
		Elem: tf_schema.TypeBool,
	}, schema.Option{DropReadOnly: false})
	want := "{_:bool}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfValueTypeIntExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeMap,
		Elem: tf_schema.TypeInt,
	}, schema.Option{DropReadOnly: false})
	want := "{_:int}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfValueTypeFloatExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeMap,
		Elem: tf_schema.TypeFloat,
	}, schema.Option{DropReadOnly: false})
	want := "{_:float}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestMapOfValueTypeStringExpr(t *testing.T) {
	got := schema.SchemaExpr(&tf_schema.Schema{
		Type: tf_schema.TypeMap,
		Elem: tf_schema.TypeString,
	}, schema.Option{DropReadOnly: true})
	want := "{_:string}"
	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}
