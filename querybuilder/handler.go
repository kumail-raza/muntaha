package querybuilder

import "reflect"

type Handler interface {
	Handle(mv reflect.Value, mt reflect.StructField, args NeoArgs) (string, error)
}

// NewHandler creates a new handler for neo4j models
func NewHandler(kind string, v interface{}) Handler {

	//TODO:: add other types
	switch kind {
	case "string":
		return NewStringHandler(v)
	case "slice":
		return NewSliceHandler(v)
	case "int64":
		return NewInt64Handler(v)
	case "ptr":
		return NewPtrHandler(v)
	}
	return nil
}
