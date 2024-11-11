package mirusort

func BubbleSort(nums []int) []int {
	N := len(nums)

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
