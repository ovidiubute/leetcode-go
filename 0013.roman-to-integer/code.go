package romantointeger

func romanToInt(s string) int {
	intValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	v := 0

	for i := 0; i < len(s); {
		c1 := string(s[i])

		if i < len(s)-1 {
			c2 := string(s[i+1])

			if intValues[c2] > intValues[c1] {
				v += intValues[c2] - intValues[c1]
				i += 2
				continue
			}
		}

		v += intValues[c1]
		i++
	}

	return v
}
