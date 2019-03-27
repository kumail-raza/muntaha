package querybuilder

import (
	"reflect"
)

type PtrHandler interface {
	Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error)
}

func NewPtrHandler(v interface{}) PtrHandler {
	return &ptrHandler{
		ModelType:  reflect.TypeOf(v),
		ModelValue: reflect.ValueOf(v),
	}
}

type ptrHandler struct {
	ModelValue reflect.Value
	ModelType  reflect.Type
}

func (p *ptrHandler) Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error) {

	return "", nil
}
