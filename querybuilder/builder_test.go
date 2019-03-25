package querybuilder

import (
	"testing"

	"github.com/minhajuddinkhan/muntaha/models"
)

const (
	duaRefVar = ModelReference("d")
	refRefVar = ModelReference("r")
	orgRefVar = ModelReference("o")
	emoRefVar = ModelReference("e")
)

func TestDuaArgModel(t *testing.T) {

	args := make(NeoArgs)
	expected := `(d:Dua{name: {duaName},title: {duaTitle},translation: {translation})`
	input := models.Dua{Title: "123", Arabic: "456", Translation: "789"}
	outputQ, args := NewDuaArgModel(input, args, duaRefVar)
	if outputQ != expected {
		t.Errorf("%s should be equal to %s", outputQ, expected)
	}
	if args["duaName"] != input.Arabic {
		t.Errorf("%s should be equal to %s", args["duaName"], input.Arabic)
	}
	if args["duaTitle"] != input.Title {
		t.Errorf("%s should be equal to %s", args["duaTitle"], input.Arabic)
	}
	if args["translation"] != input.Translation {
		t.Errorf("%s should be equal to %s", args["translation"], input.Arabic)
	}
}

func TestRefArgModel(t *testing.T) {

	input := models.Reference{Name: "123"}
	args := make(NeoArgs)
	expected := `(r:Reference{name: {refName}})`
	outputQ, args := NewRefArgModel(input, args, refRefVar)

	if outputQ != expected {
		t.Errorf("%s should be equal to %s", outputQ, expected)
	}

	if args["refName"] != input.Name {
		t.Errorf("%s should be equal to %s", args["name"], input.Name)
	}

}

func TestNewEmotionArgModel(t *testing.T) {
	input := models.Emotion{Name: "1"}
	args := make(NeoArgs)
	expected := `(e:Emotion{name: {emoName}})`

	outputQ, args := NewEmotionArgModel(input, args, emoRefVar)
	if outputQ != expected {
		t.Errorf("%s should be equal to %s", outputQ, expected)
		return
	}
	if args["emoName"] != input.Name {
		t.Errorf("%s should be equal to %s", args["emoName"], input.Name)
		return
	}
}

func TestNewOriginArgModel(t *testing.T) {
	o := models.Origin{Type: "1", References: []models.Reference{
		models.Reference{Name: "2"},
	}}
	expected := `(o:Origin{type: {orgType},references: {orgRefs}})`
	q, args := NewOriginArgModel(o, make(NeoArgs), orgRefVar)
	if q != expected {
		t.Errorf("%s should be equal to %s", q, expected)
		return
	}
	if args["orgType"] != o.Type {
		t.Errorf("%s should be equal to %s", args["orgType"], o.Type)
		return
	}
}
