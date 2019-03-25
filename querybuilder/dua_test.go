package querybuilder

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/minhajuddinkhan/muntaha/models"
)

func TestCreateRelationInRefAndDua(t *testing.T) {
	ref := models.Reference{Name: "1235"}
	d := models.Dua{Arabic: "1", Title: "2", Translation: "3"}
	outputQ, args := CreateRelationInRefAndDua(ref.Name, d)
	expected := `MATCH (r:Reference{name: {refName}}),(d:Dua{name: {duaName},title: {duaTitle},translation: {translation}) CREATE (d)-[:REFERENCED_IN]->(r)`
	if outputQ != expected {
		t.Errorf("%s should be equal to %s", outputQ, expected)
		return
	}
	if args["refName"] != ref.Name {
		t.Errorf("%s should be equal to %s", args["refName"], ref.Name)
		return
	}
}

func TestGetDuaByTitle(t *testing.T) {
	title := "123"
	expectedQuery := "MATCH (d:Dua{title: {duaTitle}}) RETURN d"

	outputQuery, args := GetDuaByTitle(title)
	if expectedQuery != outputQuery {
		t.Errorf("%s should be equal to %s", outputQuery, expectedQuery)
		return
	}
	if args["duaTitle"] != title {
		t.Errorf("%s should be equal to %s", args["duaTitle"], title)
	}
}

func TestCreateRelationInEmoAndDua(t *testing.T) {

	e := models.Emotion{Name: "happy"}
	d := models.Dua{Arabic: "123"}
	expectedQ := `MATCH (e:Emotion{name: {emoName}}),(d:Dua{name: {duaName}}) CREATE (e)-[:RELATED]->(d)`
	q, args := CreateRelationInEmoAndDua(e, d)
	if expectedQ != q {
		t.Errorf("%s should be equal to %s", q, expectedQ)
		return
	}
	if args["emoName"] != e.Name {
		t.Errorf("%s should be equal to %s", args["emoName"], e.Name)
		return
	}

	spew.Dump(q, args)
}

func TestEmptyModel(t *testing.T) {
	q, _ := NewDuaArgModel(models.Dua{}, make(NeoArgs))
	expected := `(d:Dua{})`
	if q != expected {
		t.Errorf("%s should be equal to %s", q, expected)
	}
}
