package querybuilder

import (
	"fmt"

	"github.com/minhajuddinkhan/muntaha/models"
)

// NewDuaArgModel creates a new dua argument model for neo4j
func NewDuaArgModel(dua models.Dua, args NeoArgs) (string, NeoArgs) {

	duaRaw := ""
	if dua.Arabic != "" {
		duaRaw += "name: {duaName},"
		args["duaName"] = dua.Arabic
	}
	if dua.Title != "" {
		duaRaw += "title: {duaTitle},"
		args["duaTitle"] = dua.Title
	}
	if dua.Translation != "" {
		duaRaw += "translation: {translation}"
		args["translation"] = dua.Translation
	}
	duaRaw = normalizeRawQuery(duaRaw)
	return fmt.Sprintf("(d:Dua{%s})", duaRaw), args

}

// NewRefArgModel creates a new reference model for neo4j
func NewRefArgModel(ref models.Reference, args NeoArgs) (string, NeoArgs) {

	refRaw := ""
	if ref.Name != "" {
		refRaw += `name: {refName},`
		args["refName"] = ref.Name
	}

	//TODO:: add ref numbers
	//TODO:: add ref Id

	refRaw = normalizeRawQuery(refRaw)
	return fmt.Sprintf("(r:Reference{%s})", refRaw), args

}

// NewEmotionArgModel creates a new emotion model for neo4j
func NewEmotionArgModel(emo models.Emotion, args NeoArgs) (string, NeoArgs) {
	emoRaw := ""
	if emo.Name != "" {
		emoRaw += `name: {emoName},`
		args["emoName"] = emo.Name
	}
	emoRaw = normalizeRawQuery(emoRaw)
	return fmt.Sprintf("(e:Emotion{%s})", emoRaw), args
}

// NewOriginArgModel creates new origin model for neo4j
func NewOriginArgModel(o models.Origin, args NeoArgs) (string, NeoArgs) {

	originRaw := ""
	if o.Type != "" {
		originRaw += `type: {orgType},`
		args["orgType"] = o.Type
	}
	if len(o.References) != 0 {
		originRaw += "references: {orgRefs},"
		refs := make([]interface{}, len(o.References))
		for j, ref := range o.References {
			refs[j] = ref.Name
		}
		args["orgRefs"] = refs
	}
	originRaw = normalizeRawQuery(originRaw)
	return fmt.Sprintf("(o:Origin{%s})", originRaw), args

}

func normalizeRawQuery(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[0 : len(s)-1]
}
