package apidts

import (
	"strings"
	"testing"
)

func convertThenStringize(s string) (string, error) {
	reader := strings.NewReader(s)
	var e error

	d, e := ConvertJsonToDts(reader)
	if e != nil {
		return "", e
	}

	r, e := StringizeDts(d)
	if e != nil {
		return "", e
	}

	return r, nil
}

func TestStringizeDts(t *testing.T) {

	check := func(json string, expected string) {
		actual, e := convertThenStringize(json)
		if e != nil {
			t.Errorf("'%s' must not occur error", json)
		}
		if actual != expected {
			t.Errorf("Expected '%v' but actually '%v'", expected, actual)
		}
	}

	check(`{"foo": true}`, `interface FixMe {
  foo: boolean;
}`)
	check(`[{"foo": true}, {"foo": false}]`, `interface FixMe {
  foo: boolean;
}`)
	check(`{"bar": 42}`, `interface FixMe {
  bar: number;
}`)
	check(`{"foo": [1, 2, 3]}`, `interface FixMe {
  foo: number[];
}`)
	check(`{"foo": {"piyo": "foo", "poyo": [1, 2, 3]}}`, `interface FixMe {
  foo: {
    piyo: string;
    poyo: number[];
  };
}`)

	if _, e := convertThenStringize("true"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}

	if _, e := convertThenStringize("42"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}

	if _, e := convertThenStringize("null"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}
}
