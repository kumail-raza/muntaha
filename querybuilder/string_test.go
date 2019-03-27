package querybuilder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHandler(t *testing.T) {

	values := []string{"John Doe", "123"}
	keys := []string{"SimpleModel.name", "SimpleModel.age"}
	s := SimpleModel{Name: values[0], Age: values[1]}

	queries := []string{`name: {SimpleModel.name},`, "age: {SimpleModel.age},"}
	modelType := reflect.TypeOf(s)
	modelValue := reflect.ValueOf(s)
	sh := NewStringHandler(s)

	for i := 0; i < modelValue.NumField(); i++ {
		fieldType := modelType.Field(i).Type.Kind().String()
		args := make(NeoArgs)
		switch fieldType {
		case "string":
			q, err := sh.Handle(modelValue.Field(i), modelType.Field(i), args)
			assert.Nil(t, err)
			assert.Equal(t, queries[i], q)
			assert.Equal(t, args[keys[i]], values[i])
		}
	}
}
