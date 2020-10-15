package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		_, ok := seen[nums[i]]

		if ok {
			if i-seen[nums[i]] <= k {
				return true
			}
		}
		seen[nums[i]] = i
	}

	return false
}
