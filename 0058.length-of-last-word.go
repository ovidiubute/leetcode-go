package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	fields := strings.Fields(s)

	if (len(fields) == 0) {
		return 0
	}
	
	return len(fields[len(fields) - 1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
