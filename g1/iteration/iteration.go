package iteration

import (
	"strings"
)

func Repeat(ch string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(ch)
	}
	return repeated.String()
}
