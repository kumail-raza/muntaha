package querybuilder

type SimpleModel struct {
	Name string `neo:"name,primary"`
	Age  string `neo:"age"`
}
type ModelWithArray struct {
	Type          string        `neo:"type"`
	ArrayOfModels []SimpleModel `neo:"simpleModels"`
}

type ModelWithModel struct {
	Referred SimpleModel `neo:"simpleModel"`
}
