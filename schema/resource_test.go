package schema_test

import (
	"testing"
	"reflect"
	"github.com/westelh/tfprovider-cue/schema"
	"cuelang.org/go/cue/ast"
	tf_schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceWithComments(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type: tf_schema.TypeString,
				Description: "This is a foo",
			},
		},
	})
	want := "{// This is a foo\nfoo: string}"

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}

func TestResourceWithDefault(t *testing.T) {
	var got ast.Expr = schema.ResourceExpr(&tf_schema.Resource{
		Schema: map[string]*tf_schema.Schema{
			"foo": {
				Type: tf_schema.TypeString,
				Default: "bar",
			},
		},
	})
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
	})
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
	})
	want := `{foo: string}`

	if reflect.DeepEqual(FormatExpr(&got), FormatString(want)) {
		t.Fatalf("unexpected cue: %s expected: %s", FormatExpr(&got), FormatString(want))
	}
}
