package mirusort

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		j++
	}
	return result
}

func MergeSort(nums []int) []int {
	var N = len(nums)
	if N == 1 {
		return nums
	}
	mid := N / 2
	left := make([]int, mid)
	right := make([]int, N-mid)

	for i := 0; i < N; i++ {
		if i < mid {
			left[i] = nums[i]
		} else {
			right[i-mid] = nums[i]
		}
	}
	return merge(MergeSort(left), MergeSort(right))
}
