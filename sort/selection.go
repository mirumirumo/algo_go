package mirusort

func SelectionSort(nums []int) []int {
	N := len(nums)
	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums

}
