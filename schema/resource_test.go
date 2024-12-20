package schema_test

import (
	"reflect"
	"testing"

	"cuelang.org/go/cue/ast"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/westelh/tfprovider-cue/schema"
)

func TestResourceWithComments(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type:        tf_schema.TypeString,
				Description: "This is a foo",
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := "{// This is a foo\nfoo: string}"

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestResourceWithDefault(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type:    tf_schema.TypeString,
				Default: "bar",
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := `{foo: *bar | string}`

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestResourceWithOptional(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type:     tf_schema.TypeString,
				Optional: true,
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := `{foo?: string}`

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestResourceWithRequired(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type:     tf_schema.TypeString,
				Required: true,
			},
		},
	}, schema.Option{DropReadOnly: false})
	want := `{foo: string}`

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestDropReadOnly(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type:     tf_schema.TypeString,
				Optional: false,
				Required: false,
				Computed: true,
			},
			"bar": {
				Type:     tf_schema.TypeString,
				Optional: false,
				Required: false,
				Computed: false,
			},
		},
	}, schema.Option{DropReadOnly: true})
	want := `{bar: string}`

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}
