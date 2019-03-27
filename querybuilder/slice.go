package querybuilder

import (
	"reflect"
	"strings"
)

type SliceHandler interface {
	Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error)
}

func NewSliceHandler(v interface{}) SliceHandler {
	return &sliceHandler{
		ModelType:  reflect.TypeOf(v),
		ModelValue: reflect.ValueOf(v),
	}
}

type sliceHandler struct {
	ModelValue reflect.Value
	ModelType  reflect.Type
}

func (p *sliceHandler) Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error) {
	var result []interface{}
	for j := 0; j < mv.Len(); j++ {
		slideField := reflect.TypeOf(mv).Field(j)
		switch slideField.Type.Kind().String() {
		case "ptr":
			item := reflect.TypeOf(mv.Index(j).Interface())
			var primaryFieldIndex int
			for k := 0; k < item.NumField(); k++ {
				tag := item.Field(k).Tag.Get("neo")
				if strings.Contains(tag, "primary") {
					primaryFieldIndex = k
					break
				}
			}
			x := mv.Index(j).Field(primaryFieldIndex).String()
			result = append(result, x)
		}
	}

	taggedName := mt.Tag.Get("neo")
	taggedName, err := getAttrName(taggedName)
	if err != nil {
		return "", err
	}
	queryAttrs := AttributeName(p.ModelType.Name(), taggedName)
	args[queryAttrs] = result
	return QueryFmt(taggedName, queryAttrs), nil

}
