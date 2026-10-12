package iteration

import (
	"strings"
)

func Repeat(ch string) string {
	var repeated strings.Builder
	countTimes := 5
	for i := 0; i < countTimes; i++ {
		repeated.WriteString(ch)
	}
	return repeated.String()
}
