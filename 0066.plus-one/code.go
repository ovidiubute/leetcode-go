package plusone

func plusOne(digits []int) []int {
	carry := 0
	var out = make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			out[i+1] = digits[i] + 1
		} else {
			out[i+1] = digits[i] + carry
		}

		if out[i+1] > 9 {
			carry = 1
			out[i+1] = out[i+1] % 10
		} else {
			carry = 0
		}
	}

	if carry > 0 {
		out[0] = 1
		return out
	}

	return out[1:]
}
