/*
 * @lc app=leetcode id=1330 lang=golang
 *
 * [1330] Reverse Subarray To Maximize Array Value
 *
 * https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/description/
 *
 * algorithms
 * Hard (35.16%)
 * Likes:    150
 * Dislikes: 23
 * Total Accepted:    2.2K
 * Total Submissions: 6.2K
 * Testcase Example:  '[2,3,1,5,4]'
 *
 * You are given an integer array nums. The value of this array is defined as
 * the sum of |nums[i]-nums[i+1]| for all 0 <= i < nums.length-1.
 * 
 * You are allowed to select any subarray of the given array and reverse it.
 * You can perform this operation only once.
 * 
 * Find maximum possible value of the final array.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,3,1,5,4]
 * Output: 10
 * Explanation: By reversing the subarray [3,1,5] the array becomes [2,5,1,3,4]
 * whose value is 10.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [2,4,9,24,2,1,10]
 * Output: 68
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 3*10^4
 * -10^5 <= nums[i] <= 10^5
 * 
 */

// @lc code=start
func maxValueAfterReverse(nums []int) int {
    
}
// @lc code=end
