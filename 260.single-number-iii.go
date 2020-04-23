/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
// singleNumber use xor
// https://leetcode.com/problems/single-number-iii/discuss/68900/Accepted-C%2B%2BJava-O(n)-time-O(1)-space-Easy-Solution-with-Detail-Explanations
func singleNumber(nums []int) []int {
	// XOR all number, gives us a XOR b, where a and b are the ones we wanted

	// If bit_i in (a xor b) is 1, bit_i at a and b are different.
	// Find bit_i using the low bit formula m & -m
	// Partition the numbers into two groups: one group with bit_i == 1 and the other group with bit_i == 0.
	//   detail:
	// a is in one group and b is in the other.
	// a is the only single number in its group.
	// b is also the only single number in its group.
	// XORing all numbers in a's group to get a
	// XORing all numbers in b's group to get b
	// Alternatively, XOR (a xor b) with a gets you b.
	var aXORb int = 0
	var a, b int = 0, 0
	for _, i := range nums {
		aXORb ^= i
	}
	lower_bit := aXORb & -aXORb // the lower_bit well be there
	for _, i := range nums {
		if i&lower_bit == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}
	return []int{a, b}
}

func singleNumberMap(nums []int) []int {
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
