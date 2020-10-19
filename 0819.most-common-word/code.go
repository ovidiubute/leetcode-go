package mostcommonword

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedWords := map[string]int{}

	for _, w := range banned {
		bannedWords[w] = 1
	}

	words := strings.FieldsFunc(strings.ToLower(paragraph), func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	freq := map[string]int{}

	max := -1
	maxPos := 0
	for i := 0; i < len(words); i++ {
		_, ok := bannedWords[words[i]]

		if ok {
			continue
		}

		v, ok := freq[words[i]]
		if !ok {
			freq[words[i]] = 1
		} else {
			freq[words[i]] = v + 1
		}

		if max < freq[words[i]] {
			max = freq[words[i]]
			maxPos = i
		}
	}

	return words[maxPos]
}
