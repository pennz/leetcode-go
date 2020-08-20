package utils

func quickSort(nums []string, compareFunc func(string, string) bool) {
	if len(nums) <= 1 {
		return
	}

	pivot := partition(nums, compareFunc)
	ls, rs := nums[:pivot], nums[pivot:]
	quickSort(ls, compareFunc)
	quickSort(rs, compareFunc)
}
func partition(nums []string, compareFunc func(string, string) bool) int {
	pivot, pv := 0, nums[0]
	if pv == "" {
		pivot += 1
	}
	return pivot
}
