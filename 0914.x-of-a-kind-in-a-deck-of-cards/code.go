package xofakindinadeckofcards

func hasGroupsSizeX(deck []int) bool {
	count := map[int]int{}

	for _, v := range deck {
		count[v]++
	}

	g := -1
	for _, c := range count {
		if g == -1 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}

	return g >= 2
}

func gcd(x int, y int) int {
	if x == 0 {
		return y
	}

	return gcd(y%x, x)
}
