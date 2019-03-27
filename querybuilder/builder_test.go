package querybuilder

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

const (
	duaRefVar = ModelReference("d")
	refRefVar = ModelReference("r")
	orgRefVar = ModelReference("o")
	emoRefVar = ModelReference("e")
)

func TestNewModelWithArrayInterface(t *testing.T) {

	mwa := ModelWithArray{Type: "type", ArrayOfModels: []SimpleModel{
		SimpleModel{Name: "sm"},
	}}
	var intfs []interface{}
	for _, i := range mwa.ArrayOfModels {
		intfs = append(intfs, i.Name)
	}
	args := make(NeoArgs)
	expectedQ := `(mwa:ModelWithArray{type: {ModelWithArray.type},simpleModels: {ModelWithArray.simpleModels}})`
	q, args, err := NewModel(mwa, args, "mwa")
	assert.Nil(t, err)
	assert.Equal(t, expectedQ, q)
	assert.Equal(t, args["ModelWithArray.type"], mwa.Type)
	assert.Equal(t, args["ModelWithArray.simpleModels"], intfs)
}
func TestNewModelWithStringAttributes(t *testing.T) {

	sm := SimpleModel{Name: "mySimpleModel"}

	expectedQ := `(sm:SimpleModel{name: {SimpleModel.name}})`
	args := make(NeoArgs)
	q, args, err := NewModel(sm, args, "sm")
	assert.Nil(t, err)

	assert.Equal(t, expectedQ, q)
	assert.Equal(t, sm.Name, args["SimpleModel.name"])
}

func TestNewModelWithModelInAModel(t *testing.T) {
	mwm := ModelWithModel{Referred: SimpleModel{Name: "123"}}
	args := make(NeoArgs)
	q, args, err := NewModel(mwm, args, "mwm")
	assert.Nil(t, err)
	spew.Dump(q, args)

}
