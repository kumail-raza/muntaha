package querybuilder

import (
	"fmt"
	"strings"
)

func normalizeRawQuery(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[0 : len(s)-1]
}

func getAttrName(s string) (string, error) {

	var unknowns []string
	arr := strings.Split(s, ",")
	for _, el := range arr {
		if _, ok := keywords[el]; !ok {
			unknowns = append(unknowns, el)
		}
	}
	if len(unknowns) > 1 {
		return "", fmt.Errorf("ambigious query attribute")
	}
	return unknowns[0], nil
}
