package lengthoflastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	fields := strings.Fields(s)

	if len(fields) == 0 {
		return 0
	}

	return len(fields[len(fields)-1])
}
