package main

import "gitlab.com/MrCue/leetcode-go/utils"

/*
 * @lc app=leetcode id=421 lang=golang
 *
 * [421] Maximum XOR of Two Numbers in an Array
 *
 * https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/description/
 *
 * algorithms
 * Medium (53.52%)
 * Likes:    1240
 * Dislikes: 170
 * Total Accepted:    52.2K
 * Total Submissions: 97.5K
 * Testcase Example:  '[3,10,5,25,2,8]'
 *
 * Given a non-empty array of numbers, a0, a1, a2, … , an-1, where 0 ≤ ai <
 * 2^31.
 *
 * Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.
 *
 * Could you do this in O(n) runtime?
 *
 * Example:
 *
 *
 * Input: [3, 10, 5, 25, 2, 8]
 *
 * Output: 28
 *
 * Explanation: The maximum result is 5 ^ 25 = 28.
 *
 *
 *
 *
 */

// @lc code=start
func findXORpair(a int, t *utils.BinaryTrie) int {
	mask := 1 << 31
	guard := t.Root
	for i := 0; i < 32; i, mask = i+1, mask>>1 {
		bit := 1
		if a&mask == 0 {
			bit = 0
		}

		if guard.Children[1-bit] != nil {
			guard = guard.Children[1-bit]
		} else {
			guard = guard.Children[bit]
		}
	}
	if !guard.IsWordEnd {
		panic("Should be at end now")
	}
	return guard.Value
}
func findMaximumXOR(nums []int) int {
	a := utils.InitBinaryTrie()
	for _, n := range nums {
		a.Insert(n)
	}

	var maxXOR int
	for _, n := range nums {
		p := findXORpair(n, a)
		x := p ^ n
		if x > maxXOR {
			maxXOR = x
		}
	}

	return maxXOR
}

// @lc code=end
