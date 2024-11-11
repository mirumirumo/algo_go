package mirusort

func InsertionSort(nums []int) []int {
	N := len(nums)
	for i := 1; i < N; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = tmp
	}
	return nums
}
