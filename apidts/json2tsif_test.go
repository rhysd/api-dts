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

	var e error

	b, e := ConvertJsonToDts(reader("true"))
	if e != nil {
		t.Errorf("b must not occur an error: %v", e)
	}
	if b.token != TypeBoolean {
		t.Errorf("Type must be boolean but %v", b.token)
	}

	n, e := ConvertJsonToDts(reader("42"))
	if e != nil {
		t.Errorf("n must not occur an error: %v", e)
	}
	if n.token != TypeNumber {
		t.Errorf("Type must be number but %v", n.token)
	}

	s, e := ConvertJsonToDts(reader(`"foo"`))
	if e != nil {
		t.Errorf("s must not occur an error: %v", e)
	}
	if s.token != TypeString {
		t.Errorf("Type must be string but %v", s.token)
	}

	n2, e := ConvertJsonToDts(reader("null"))
	if e != nil {
		t.Errorf("n2 must not occur an error: %v", e)
	}
	if n2.token != TypeNull {
		t.Errorf("Type must be null but %v", n2.token)
	}

	a, e := ConvertJsonToDts(reader("[1, 2, 3]"))
	if e != nil {
		t.Errorf("a must not occur an error: %v", e)
	}
	if a.token != TypeArray {
		t.Errorf("Type must be array but %v", a.token)
	}
	if a.elem_type == nil {
		t.Errorf("'elem_type' must not be nil for array")
	}
	if a.elem_type.token != TypeNumber {
		t.Errorf("'elem_type' must be element type number but actually %v", a.elem_type.token)
	}

	if r, _ := ConvertJsonToDts(reader("[]")); r.elem_type.token != TypeNull {
		t.Errorf("'elem_type' of empty array must be null")
	}

	o, e := ConvertJsonToDts(reader(`{"foo": 1, "bar": 2}`))
	if e != nil {
		t.Errorf("o must not occur an error: %v", e)
	}
	if o.token != TypeObject {
		t.Errorf("Type must be object but %v", o.token)
	}
	if o.fields == nil {
		t.Errorf("'fields' must not be nil for object")
	}
	if o.fields["foo"].token != TypeNumber {
		t.Errorf("'foo' field must have TypeNumber token but %v", o.fields["foo"].token)
	}

	o2, e := ConvertJsonToDts(reader(`{"foo": {"poyo": true}, "bar": {"puyo": false}}`))
	if e != nil {
		t.Errorf("o2 must not occur an error: %v", e)
	}
	if o2.fields["foo"].token != TypeObject {
		t.Errorf("'foo' field must have TypeObject token but %v", o2.fields["foo"].token)
	}
	if o2.fields["foo"].fields["poyo"].token != TypeBoolean {
		t.Errorf("'foo' field's child must have field 'poyo' and its token TypeBoolean but %v", o2.fields["foo"].fields["foo"].token)
	}
}
