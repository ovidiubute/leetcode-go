package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 || (x > 9 && x%10 == 0) {
		return false
	}

	r := 0
	for r < x {
		r = r*10 + x%10
		x /= 10
	}

	return x == r || x == r/10
}
