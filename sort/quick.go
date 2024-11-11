package mirusort

func partition(nums []int, low, high int) int {
	i := low
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func quickSortRecursion(nums []int, low, high int) []int {
	if low < high {
		partitionIndex := partition(nums, low, high)
		quickSortRecursion(nums, low, partitionIndex-1)
		quickSortRecursion(nums, partitionIndex+1, high)
	}
	return nums
}

func QuickSort(nums []int) []int {
	return quickSortRecursion(nums, 0, len(nums)-1)
}
