func findMaxConsecutiveOnes(nums []int) int {
	maximum := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count += 1
		} else {
			if count > maximum {
				maximum = count
			}
			count = 0
		}
	}
	if count > maximum {
		maximum = count
	}
	return maximum
}