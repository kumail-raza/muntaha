package repo

import (
	"fmt"

	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/querybuilder"
)

const getAllDua = `MATCH(d:Dua) RETURN d`
const createDua = `CREATE(d:Dua {name: {name},title: {title}, references: {references}})`

// GetDuaByTitle creates query and args for fetching dua by title
func GetDuaByTitle(title string) (string, querybuilder.NeoArgs, error) {
	args := make(map[string]interface{})
	duaRef := "d"
	q, args, err := querybuilder.NewModel(models.Dua{Title: title}, args, "d")
	return fmt.Sprintf("MATCH %s RETURN %s", q, duaRef), args, err
}

// GetAllDua creates query and args for fetching all duas
func GetAllDua() (string, querybuilder.NeoArgs, error) {
	return querybuilder.NewModel(models.Dua{}, nil, "d")
}

// CreateRelationInRefAndDua creates relationship dua -> reference
func CreateRelationInRefAndDua(refName string, dua models.Dua) (string, querybuilder.NeoArgs, error) {

	relation := "REFERENCED_IN"
	return querybuilder.CreateRelationAToB(relation, dua, models.Reference{Name: refName})
}

// CreateRelationInEmoAndDua creates relation emotion -> dua
func CreateRelationInEmoAndDua(emo models.Emotion, dua models.Dua) (string, querybuilder.NeoArgs, error) {

	relation := "RELATED"
	return querybuilder.CreateRelationAToB(relation, emo, dua)

}

// CreateRelationInOriginAndDua creates relation dua -> emo
func CreateRelationInOriginAndDua(o models.Origin, dua models.Dua) (string, querybuilder.NeoArgs, error) {

	relation := "REFERENCED_IN"
	return querybuilder.CreateRelationAToB(relation, dua, o)
}

// CreateDua creates query and args for creating a dua
func CreateDua(d models.Dua, o models.Origin) (string, querybuilder.NeoArgs) {

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
