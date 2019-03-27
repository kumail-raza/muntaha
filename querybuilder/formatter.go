package querybuilder

import "fmt"

func AttributeName(modelName, attrName string) string {
	return fmt.Sprintf("%s%s", modelName, attrName)
}

func QueryFmt(modelName, queryAttrs string) string {
	return fmt.Sprintf("%s: {%s},", modelName, queryAttrs)
}
