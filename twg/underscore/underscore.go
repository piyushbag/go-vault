package underscore

import "strings"

func Camel(str string) string {
	var result strings.Builder
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result.WriteString("_")
		}
		result.WriteRune(char)
	}
	return strings.ToLower(result.String())
}
