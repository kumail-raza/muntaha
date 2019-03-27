package querybuilder

import (
	"reflect"
)

func NewInt64Handler(v interface{}) Handler {
	return &int64Hanlder{
		ModelType:  reflect.TypeOf(v),
		ModelValue: reflect.ValueOf(v),
	}
}

type int64Hanlder struct {
	ModelValue reflect.Value
	ModelType  reflect.Type
}

func (s *int64Hanlder) Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error) {

	fieldVal := mv.Interface().(int64)
	if fieldVal == 0 {
		return "", nil
	}
	taggedName := mt.Tag.Get("neo")
	taggedName, err := getAttrName(taggedName)
	if err != nil {
		return "", err
	}

	queryAttrs := AttributeName(s.ModelType.Name(), taggedName)
	args[queryAttrs] = fieldVal
	return QueryFmt(taggedName, queryAttrs), nil

}
