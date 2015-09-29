package apidts

import (
	"strings"
	"testing"
)

func TestNewTypeScriptDef(t *testing.T) {
	d := NewTypeScriptDef(TypeUnknown)
	if d.token != TypeUnknown {
		t.Errorf("Want %v but %v", TypeUnknown, d)
	}
	if d.elem_type != nil {
		t.Errorf("'elem_type' must be nil on init but %v", d.elem_type)
	}
	if d.fields != nil {
		t.Errorf("'fields' must be nil on init but %v", d.fields)
	}
}

func TestConvertJsonToDts(t *testing.T) {
	reader := func(s string) *strings.Reader {
		return strings.NewReader(s)
	}

	b := ConvertJsonToDts(reader("true"))
	if b.token != TypeBoolean {
		t.Errorf("Type must be boolean but %v", b.token)
	}

	n := ConvertJsonToDts(reader("42"))
	if n.token != TypeNumber {
		t.Errorf("Type must be number but %v", n.token)
	}

	s := ConvertJsonToDts(reader(`"foo"`))
	if s.token != TypeString {
		t.Errorf("Type must be string but %v", s.token)
	}

	n2 := ConvertJsonToDts(reader("null"))
	if n2.token != TypeNull {
		t.Errorf("Type must be null but %v", n2.token)
	}

	a := ConvertJsonToDts(reader("[1, 2, 3]"))
	if a.token != TypeArray {
		t.Errorf("Type must be array but %v", a.token)
	}
	if a.elem_type == nil {
		t.Errorf("'elem_type' must not be nil for array")
	}
	if a.elem_type.token != TypeNumber {
		t.Errorf("'elem_type' must be element type number but actually %v", a.elem_type.token)
	}

	if ConvertJsonToDts(reader("[]")).elem_type.token != TypeNull {
		t.Errorf("'elem_type' of empty array must be null")
	}

	o := ConvertJsonToDts(reader(`{"foo": 1, "bar": 2}`))
	if o.token != TypeObject {
		t.Errorf("Type must be object but %v", o.token)
	}
	if o.fields == nil {
		t.Errorf("'fields' must not be nil for object")
	}
	if o.fields["foo"].token != TypeNumber {
		t.Errorf("'foo' field must have TypeNumber token but %v", o.fields["foo"].token)
	}

	o2 := ConvertJsonToDts(reader(`{"foo": {"poyo": true}, "bar": {"puyo": false}}`))
	if o2.fields["foo"].token != TypeObject {
		t.Errorf("'foo' field must have TypeObject token but %v", o2.fields["foo"].token)
	}
	if o2.fields["foo"].fields["poyo"].token != TypeBoolean {
		t.Errorf("'foo' field's child must have field 'poyo' and its token TypeBoolean but %v", o2.fields["foo"].fields["foo"].token)
	}
}
