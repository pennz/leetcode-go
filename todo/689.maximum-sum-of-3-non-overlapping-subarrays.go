package main


/*
 * @lc app=leetcode id=689 lang=golang
 *
 * [689] Maximum Sum of 3 Non-Overlapping Subarrays
 *
 * https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/description/
 *
 * algorithms
 * Hard (46.31%)
 * Likes:    972
 * Dislikes: 57
 * Total Accepted:    43.8K
 * Total Submissions: 94.5K
 * Testcase Example:  '[1,2,1,2,6,7,5,1]\n2'
 *
 * In a given array nums of positive integers, find three non-overlapping
 * subarrays with maximum sum.
 * 
 * Each subarray will be of size k, and we want to maximize the sum of all 3*k
 * entries.
 * 
 * Return the result as a list of indices representing the starting position of
 * each interval (0-indexed). If there are multiple answers, return the
 * lexicographically smallest one.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,1,2,6,7,5,1], 2
 * Output: [0, 3, 5]
 * Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting
 * indices [0, 3, 5].
 * We could have also taken [2, 1], but an answer of [1, 3, 5] would be
 * lexicographically larger.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * nums.length will be between 1 and 20000.
 * nums[i] will be between 1 and 65535.
 * k will be between 1 and floor(nums.length / 3).
 * 
 * 
 * 
 * 
 */

// @lc code=start
func maxSumOfThreeSubarrays(nums []int, k int) []int {
    
}
// @lc code=end
