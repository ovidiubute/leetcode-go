package missingnumber

func missingNumber(nums []int) int {
	seen := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		seen[i] = 0
	}

	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = 1
	}

	for i := 0; i < len(seen); i++ {
		if seen[i] == 0 {
			return i
		}
	}

	return -1
}
