package querybuilder

import (
	"fmt"

	"github.com/minhajuddinkhan/muntaha/models"
)

const duaByTitle = `MATCH(d:Dua{ title: {title} }) RETURN d`
const createDua = `CREATE(d:Dua {name: {name},title: {title}, references: {references}})`
const getAllDua = `MATCH(d:Dua) RETURN d`

// NeoArgs neo4j go argument style
type NeoArgs map[string]interface{}

// CreateDua creates query and args for creating a dua
func CreateDua(d models.Dua, o models.Origin) (string, NeoArgs) {

	refs := make([]interface{}, len(o.References))
	args := make(map[string]interface{})
	for j, ref := range o.References {
		refs[j] = fmt.Sprintf("%s %s", ref.Name, ref.RefNumber)
	}
	args["references"] = refs
	args["title"] = d.Title
	args["name"] = d.Arabic

	return createDua, args
}

// GetDuaByTitle creates query and args for fetching dua by title
func GetDuaByTitle(title string) (string, NeoArgs) {
	args := make(map[string]interface{})
	duaRef := "d"
	q, args := NewDuaArgModel(models.Dua{Title: title}, args, "d")
	return fmt.Sprintf("MATCH %s RETURN %s", q, duaRef), args
}

// GetAllDua creates query and args for fetching all duas
func GetAllDua() (string, NeoArgs) {
	return getAllDua, nil
}

// CreateRelationInRefAndDua creates relationship dua -> reference
func CreateRelationInRefAndDua(refName string, dua models.Dua) (string, NeoArgs) {

	args := make(NeoArgs)
	duaRef := ModelReference("d")
	refRef := ModelReference("r")
	duaModel, args := NewDuaArgModel(dua, args, duaRef)
	refModel, args := NewRefArgModel(models.Reference{Name: refName}, args, refRef)

	q := fmt.Sprintf(`MATCH %s,%s CREATE (%s)-[:REFERENCED_IN]->(%s)`, refModel, duaModel, duaRef, refRef)
	return q, args

}

// CreateRelationInEmoAndDua creates relation emotion -> dua
func CreateRelationInEmoAndDua(emo models.Emotion, dua models.Dua) (string, NeoArgs) {

	args := make(NeoArgs)
	duaRef := ModelReference("d")
	emoRef := ModelReference("e")
	emoModel, args := NewEmotionArgModel(emo, args, emoRef)
	duaModel, args := NewDuaArgModel(dua, args, duaRef)
	q := fmt.Sprintf("MATCH %s,%s CREATE (%s)-[:RELATED]->(%s)", emoModel, duaModel, emoRef, duaRef)
	return q, args
}

// CreateRelationInOriginAndDua creates relation dua -> emo
func CreateRelationInOriginAndDua(o models.Origin, dua models.Dua) (string, NeoArgs) {

	args := make(NeoArgs)
	duaRef := ModelReference("d")
	orgRef := ModelReference("o")
	orgModel, args := NewOriginArgModel(o, args, orgRef)
	duaModel, args := NewDuaArgModel(dua, args, duaRef)
	q := fmt.Sprintf(`MATCH %s,%s CREATE (%s)-[:REFERENCED_IN]->(%s)`, orgModel, duaModel, duaRef, orgRef)
	return q, args
}
