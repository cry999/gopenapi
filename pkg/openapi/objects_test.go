package openapi

import (
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func Test(t *testing.T) {
	{
		r := strings.NewReader(`
$ref: reference
`)

		var got SchemaOrRef
		if err := yaml.NewDecoder(r).Decode(&got); err != nil {
			t.Errorf("Decode() error = %v", err)
			return
		}
		want := SchemaOrRef{
			Schema: Schema{},
			Reference: Reference{
				Ref: "reference",
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Decode() = %v, want %v", got, want)
		}
	}
	{
		r := strings.NewReader(`
type: string
format: email
`)

		var got SchemaOrRef
		if err := yaml.NewDecoder(r).Decode(&got); err != nil {
			t.Errorf("Decode() error = %v", err)
			return
		}
		want := SchemaOrRef{
			Schema:    Schema{Type: "string", Format: "email"},
			Reference: Reference{},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Decode() = %v, want %v", got, want)
		}
	}
}
