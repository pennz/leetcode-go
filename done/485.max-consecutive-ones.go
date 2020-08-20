package main

/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 *
 * https://leetcode.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (54.59%)
 * Likes:    733
 * Dislikes: 369
 * Total Accepted:    252.4K
 * Total Submissions: 462.5K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * Given a binary array, find the maximum number of consecutive 1s in this
 * array.
 *
 * Example 1:
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive
 * 1s.
 * ⁠   The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */

// @lc code=start
//  ✔ Accepted
//  ✔ 41/41 cases passed (40 ms)
//  ✔ Your runtime beats 77.05 % of golang submissions
//  ✔ Your memory usage beats 77.25 % of golang submissions (6.4 MB)
//  ✔ Accepted
//  ✔ 41/41 cases passed (40 ms)
//  ✔ Your runtime beats 77.05 % of golang submissions
//  ✔ Your memory usage beats 96.52 % of golang submissions (6.4 MB)
func findMaxConsecutiveOnes(nums []int) (cons int) {
	cons = 0
	foundC := 0

	for i := range nums {
		if nums[i] == 1 {
			foundC++
		} else {
			if foundC > cons {
				cons = foundC
			}
			foundC = 0
		}
	}

	if foundC > cons {
		cons = foundC
	} // coner case [1]
	return
}

// @lc code=end
