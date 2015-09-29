package apidts

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TypeToken uint

const (
	TypeUnknown TypeToken = iota
	TypeBoolean
	TypeNumber
	TypeString
	TypeArray
	TypeObject
	TypeNull
)

type TypeScriptDef struct {
	token     TypeToken
	elem_type *TypeScriptDef
	fields    map[string]*TypeScriptDef
}

func NewTypeScriptDef(t TypeToken) *TypeScriptDef {
	return &TypeScriptDef{
		t,
		nil,
		nil,
	}
}

func convertToArrayDef(a []interface{}) *TypeScriptDef {
	def := NewTypeScriptDef(TypeArray)
	if len(a) == 0 {
		elem := NewTypeScriptDef(TypeNull)
		def.elem_type = elem
		return def
	}

	// Note: Should check all array element types are the same
	def.elem_type = convertToDef(a[0])
	return def
}

func convertToObjDef(m map[string]interface{}) *TypeScriptDef {
	def := NewTypeScriptDef(TypeObject)
	def.fields = make(map[string]*TypeScriptDef, len(m))

	for n, t := range m {
		def.fields[n] = convertToDef(t)
	}
	return def
}

func convertToDef(val interface{}) *TypeScriptDef {
	switch t := val.(type) {
	case bool:
		return NewTypeScriptDef(TypeBoolean)
	case float64:
		return NewTypeScriptDef(TypeNumber)
	case string:
		return NewTypeScriptDef(TypeString)
	case []interface{}:
		return convertToArrayDef(t)
	case map[string]interface{}:
		return convertToObjDef(t)
	default:
		return NewTypeScriptDef(TypeNull)
	}
}

func ConvertJsonToDts(input io.Reader) *TypeScriptDef {
	var decoded interface{}
	if err := json.NewDecoder(input).Decode(&decoded); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return convertToDef(decoded)
}
