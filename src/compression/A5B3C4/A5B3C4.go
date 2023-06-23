package A5B3C4

import (
	"fmt"
)

func Compress(in string) string {
	if len(in) == 0 {
		return ""
	}

	var previous rune
	var result string
	var count int

	for i, item := range in {
		if i == 0 {
			previous = item
			continue
		}

		if item == previous {
			count++
			continue
		}

		if item != previous || i == len(in)-1 {
			result += fmt.Sprintf("%c%d", previous, count+1)
		}

		count = 0
		previous = item
	}

	result += fmt.Sprintf("%c%d", previous, count+1)

	return result
}
