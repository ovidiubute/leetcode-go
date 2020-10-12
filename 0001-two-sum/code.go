package twosum

func twoSum(nums []int, target int) []int {
	db := map[int]int{}

	for i, n := range nums {
		db[n] = i
	}

	for i, n := range nums {
		j, ok := db[target-n]

		if ok && i != j {
			return []int{i, j}
		}
	}

	return []int{0, 0}
}
