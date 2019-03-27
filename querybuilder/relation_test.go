package querybuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRelationAToB(t *testing.T) {

	s := SimpleModel{Name: "John Doe", Age: "18"}
	smRef := ModelWithArray{Type: "someType", ArrayOfModels: []SimpleModel{
		s,
	}}
	expectedQ := `MATCH (m1:SimpleModel{name: {SimpleModel.name},age: {SimpleModel.age}}),(m2:ModelWithArray{type: {ModelWithArray.type},simpleModels: {ModelWithArray.simpleModels}}) CREATE (m1)-[:HAS_A]->(m2)`
	q, _, err := CreateRelationAToB("HAS_A", s, smRef)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectedQ, q)
}
