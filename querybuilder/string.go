package querybuilder

import (
	"reflect"
)

func NewStringHandler(v interface{}) Handler {
	return &stringHanlder{
		ModelType:  reflect.TypeOf(v),
		ModelValue: reflect.ValueOf(v),
	}
}

type stringHanlder struct {
	ModelValue reflect.Value
	ModelType  reflect.Type
}

func (s *stringHanlder) Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error) {

	fieldVal := mv.Interface().(string)
	if fieldVal == "" {
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
