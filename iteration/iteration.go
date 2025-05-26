package iteration

import "strings"

func Repeat(inp string, times int) string {
	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(inp)
	}

	return repeated.String()
}
