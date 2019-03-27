package querybuilder

import (
	"fmt"
	"reflect"
)

var (
	keywords = map[string]bool{
		"omit":    true,
		"primary": true,
	}
)

type ModelReference string

// NeoArgs neo4j go argument style
type NeoArgs map[string]interface{}

func NewModel(model interface{}, args NeoArgs, refVar ModelReference) (string, NeoArgs, error) {

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	query := ""

	for i := 0; i < modelValue.NumField(); i++ {

		fieldType := modelType.Field(i).Type.Kind().String()
		handler := NewHandler(fieldType, model)
		if handler == nil {
			return "", nil, fmt.Errorf("cannot evaluate handler for type %s", fieldType)
		}
		q, err := handler.Handle(modelValue.Field(i), modelType.Field(i), args)
		if err != nil {
			return "", nil, err
		}
		query += q
	}
	query = normalizeRawQuery(query)
	return fmt.Sprintf("(%s:%s{%s})", refVar, modelType.Name(), query), args, nil

}
