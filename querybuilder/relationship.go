package querybuilder

import (
	"fmt"
)

// CreateRelationAToB return query and args to create a relation from A model to B model
func CreateRelationAToB(relTag string, A interface{}, B interface{}) (string, NeoArgs, error) {

	args := make(NeoArgs)
	refM1 := ModelReference("m1")
	refM2 := ModelReference("m2")
	m1, args, err := NewModel(A, args, refM1)
	if err != nil {
		return "", nil, err
	}
	m2, args, err := NewModel(B, args, refM2)
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf(`MATCH %s,%s CREATE (%s)-[:%s]->(%s)`, m1, m2, refM1, relTag, refM2), args, nil

}
