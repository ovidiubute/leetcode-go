package containsduplicate

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		val, ok := seen[nums[i]]

		if val == true {
			return true
		}

		if ok == false || val == false {
			seen[nums[i]] = true
		}
	}

	return false
}
