package main
/*
 * @lc app=leetcode id=1288 lang=golang
 *
 * [1288] Remove Covered Intervals
 *
 * https://leetcode.com/problems/remove-covered-intervals/description/
 *
 * algorithms
 * Medium (57.95%)
 * Likes:    182
 * Dislikes: 13
 * Total Accepted:    11.1K
 * Total Submissions: 19.1K
 * Testcase Example:  '[[1,4],[3,6],[2,8]]'
 *
 * Given a list of intervals, remove all intervals that are covered by another
 * interval in the list. Interval [a,b) is covered by interval [c,d) if and
 * only if c <= a and b <= d.
 * 
 * After doing so, return the number of remaining intervals.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: intervals = [[1,4],[3,6],[2,8]]
 * Output: 2
 * Explanation: Interval [3,6] is covered by [2,8], therefore it is
 * removed.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= intervals.length <= 1000
 * 0 <= intervals[i][0] < intervals[i][1] <= 10^5
 * intervals[i] != intervals[j] for all i != j
 * 
 * 
 */

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
    
}
// @lc code=end
