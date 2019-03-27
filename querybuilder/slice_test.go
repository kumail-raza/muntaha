package querybuilder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceHandler(t *testing.T) {

	s := ModelWithArray{
		ArrayOfModels: []SimpleModel{
			SimpleModel{Name: "John"},
		},
	}
	queries := []string{"simpleModels: {ModelWithArray.simpleModels},"}
	modelType := reflect.TypeOf(s)
	modelValue := reflect.ValueOf(s)

	for i := 0; i < modelValue.NumField(); i++ {
		fieldType := modelType.Field(i).Type.Kind().String()
		args := make(NeoArgs)
		switch fieldType {
		case "slice":
			sh := NewSliceHandler(s)
			q, err := sh.Handle(modelValue.Field(i), modelType.Field(i), args)
			assert.Nil(t, err)
			assert.Equal(t, q, queries[0])
		}

	}

}
