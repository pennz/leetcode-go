package main


/*
 * @lc app=leetcode id=1425 lang=golang
 *
 * [1425] Constrained Subsequence Sum
 *
 * https://leetcode.com/problems/constrained-subsequence-sum/description/
 *
 * algorithms
 * Hard (44.06%)
 * Likes:    309
 * Dislikes: 14
 * Total Accepted:    7.9K
 * Total Submissions: 17.9K
 * Testcase Example:  '[10,2,-10,5,20]\n2'
 *
 * Given an integer array nums and an integer k, return the maximum sum of a
 * non-empty subsequence of that array such that for every two consecutive
 * integers in the subsequence, nums[i] and nums[j], where i < j, the condition
 * j - i <= k is satisfied.
 * 
 * A subsequence of an array is obtained by deleting some number of elements
 * (can be zero) from the array, leaving the remaining elements in their
 * original order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [10,2,-10,5,20], k = 2
 * Output: 37
 * Explanation: The subsequence is [10, 2, 5, 20].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [-1,-2,-3], k = 1
 * Output: -1
 * Explanation: The subsequence must be non-empty, so we choose the largest
 * number.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [10,-2,-10,-5,20], k = 2
 * Output: 23
 * Explanation: The subsequence is [10, -2, -5, 20].
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 
 * 
 */

// @lc code=start
func constrainedSubsetSum(nums []int, k int) int {
    
}
// @lc code=end
