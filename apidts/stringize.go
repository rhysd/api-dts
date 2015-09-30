package apidts

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type StringizeError struct {
	token TypeToken
}

func (e *StringizeError) Error() string {
	return fmt.Sprintf("Invalid type token: Input must be object or array of object but token was %s", e.token)
}

func camelize(str string) string {
	buf := new(bytes.Buffer)
	heading := true
	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			if heading {
				heading = false
				buf.WriteRune(unicode.ToUpper(c))
			} else {
				buf.WriteRune(c)
			}
		} else {
			heading = true
		}
	}
	return buf.String()
}

// TODO: Add indent string customization
type DtsStringizer struct {
	counter uint
	indent  uint
	buffer  bytes.Buffer
}

func NewDtsStringizer() DtsStringizer {
	return DtsStringizer{0, 0, bytes.Buffer{}}
}

func (s *DtsStringizer) write(str string) {
	s.buffer.WriteString(str)
}

func (s *DtsStringizer) writeIndent() {
	for i := uint(0); i < s.indent; i++ {
		s.write("  ")
	}
}

func (s *DtsStringizer) visit(dts *TypeScriptDef) {
	switch dts.token {
	case TypeBoolean:
		s.write(" boolean")
	case TypeNumber:
		s.write(" number")
	case TypeString:
		s.write(" string")
	case TypeArray:
		s.visit(dts.elem_type)
		s.write("[]")
	case TypeObject:
		s.write(" {\n")
		s.indent += 1
		for n, f := range dts.fields {
			s.writeIndent()
			s.write(n)
			s.write(":")
			s.visit(f)
			s.write(";\n")
		}
		s.indent -= 1
		s.writeIndent()
		s.write("}")
	case TypeNull:
		s.write(" any")
	case TypeUnknown:
		panic("Unknown type token")
	default:
		panic(fmt.Sprintf("Invalid type token: %d", dts.token))
	}
}

func (s *DtsStringizer) Stringize(dts *TypeScriptDef, hint string) (string, error) {
	if dts.token == TypeArray {
		return s.Stringize(dts.elem_type, hint)
	}

	if dts.token != TypeObject {
		return "", &StringizeError{dts.token}
	}

	s.writeIndent()
	s.write("interface ")
	if hint != "" {
		idx := strings.IndexRune(hint, '.')
		if idx == -1 {
			s.write(camelize(hint))
		} else {
			s.write(camelize(hint[:idx]))
		}
	} else {
		s.write("FixMe")
	}
	s.visit(dts)
	return s.buffer.String(), nil
}

func StringizeDts(dts *TypeScriptDef, hint string) (string, error) {
	s := NewDtsStringizer()
	r, err := (&s).Stringize(dts, hint)
	if err != nil {
		return "", err
	}
	return r, nil
}
