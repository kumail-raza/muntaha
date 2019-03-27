package querybuilder

import "testing"

func TestNewHandler(t *testing.T) {

	x := NewHandler("string", SimpleModel{Age: "123"})
	if x == nil {
		t.Fail()
	}
}
