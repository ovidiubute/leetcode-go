package searchinsertposition

func searchInsert(nums []int, target int) int {
	for i, x := range nums {
		if x >= target {
			return i
		}
	}

	return len(nums)
}
