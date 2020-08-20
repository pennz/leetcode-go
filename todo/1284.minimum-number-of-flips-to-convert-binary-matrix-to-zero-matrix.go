package main


/*
 * @lc app=leetcode id=1284 lang=golang
 *
 * [1284] Minimum Number of Flips to Convert Binary Matrix to Zero Matrix
 *
 * https://leetcode.com/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix/description/
 *
 * algorithms
 * Hard (69.47%)
 * Likes:    188
 * Dislikes: 28
 * Total Accepted:    8.5K
 * Total Submissions: 12.2K
 * Testcase Example:  '[[0,0],[0,1]]\r'
 *
 * Given a m x n binary matrix mat. In one step, you can choose one cell and
 * flip it and all the four neighbours of it if they exist (Flip is changing 1
 * to 0 and 0 to 1). A pair of cells are called neighboors if they share one
 * edge.
 * 
 * Return the minimum number of steps required to convert mat to a zero matrix
 * or -1 if you cannot.
 * 
 * Binary matrix is a matrix with all cells equal to 0 or 1 only.
 * 
 * Zero matrix is a matrix with all cells equal to 0.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: mat = [[0,0],[0,1]]
 * Output: 3
 * Explanation: One possible solution is to flip (1, 0) then (0, 1) and finally
 * (1, 1) as shown.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: mat = [[0]]
 * Output: 0
 * Explanation: Given matrix is a zero matrix. We don't need to change it.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: mat = [[1,1,1],[1,0,1],[0,0,0]]
 * Output: 6
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: mat = [[1,0,0],[1,0,0]]
 * Output: -1
 * Explanation: Given matrix can't be a zero matrix
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * m == mat.length
 * n == mat[0].length
 * 1 <= m <= 3
 * 1 <= n <= 3
 * mat[i][j] is 0 or 1.
 * 
 * 
 */

// @lc code=start
func minFlips(mat [][]int) int {
    
}
// @lc code=end
