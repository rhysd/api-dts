package apidts

import (
	"bytes"
	"fmt"
	"os"
)

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

func (s *DtsStringizer) Stringize(dts *TypeScriptDef) string {
	if dts.token == TypeArray {
		return s.Stringize(dts.elem_type)
	}

	if dts.token != TypeObject {
		fmt.Fprintln(os.Stderr, "Input must be object or array of object.")
		os.Exit(1)
	}

	s.writeIndent()
	s.write("interface FixMe ")
	s.visit(dts)
	return s.buffer.String()
}

func StringizeDts(dts *TypeScriptDef) string {
	s := NewDtsStringizer()
	return (&s).Stringize(dts)
}