package iteration

import "strings"

// Repeat takes a charracter and repeated a certain amount of time.
func Repeat(character string) string {
	var repeated strings.Builder
	for range 5 {
		repeated.WriteString(character)
	}
	return repeated.String()
}
