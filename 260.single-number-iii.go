/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
// singleNumber use map, so time complexity is O(n),
// for space, is is linear too
func singleNumber(nums []int) []int {
	var numMap map[int]bool = make(map[int]bool)
	var foundOnce []int = make([]int, 0)
	for _, i := range nums {
		if _, ok := numMap[i]; ok {
			numMap[i] = true // exist more than once
		} else {
			numMap[i] = false
		}
	}
	for k, v := range numMap {
		//fmt.Println(k, v)
		if !v {
			foundOnce = append(foundOnce, k)
		}
	}
	return foundOnce
}

// @lc code=end
