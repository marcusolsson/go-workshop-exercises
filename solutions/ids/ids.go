package ids

import "strings"

func GenerateID(str string) string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return strings.Join(words, "-")
}
