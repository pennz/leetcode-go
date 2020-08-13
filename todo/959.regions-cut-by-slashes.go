/*
 * @lc app=leetcode id=959 lang=golang
 *
 * [959] Regions Cut By Slashes
 *
 * https://leetcode.com/problems/regions-cut-by-slashes/description/
 *
 * algorithms
 * Medium (66.21%)
 * Likes:    852
 * Dislikes: 169
 * Total Accepted:    19.8K
 * Total Submissions: 29.8K
 * Testcase Example:  '[" /","/ "]'
 *
 * In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a
 * /, \, or blank space.  These characters divide the square into contiguous
 * regions.
 * 
 * (Note that backslash characters are escaped, so a \ is represented as
 * "\\".)
 * 
 * Return the number of regions.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * [
 * " /",
 * "/ "
 * ]
 * Output: 2
 * Explanation: The 2x2 grid is as follows:
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * [
 * " /",
 * "  "
 * ]
 * Output: 1
 * Explanation: The 2x2 grid is as follows:
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * [
 * "\\/",
 * "/\\"
 * ]
 * Output: 4
 * Explanation: (Recall that because \ characters are escaped, "\\/" refers to
 * \/, and "/\\" refers to /\.)
 * The 2x2 grid is as follows:
 * 
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * [
 * "/\\",
 * "\\/"
 * ]
 * Output: 5
 * Explanation: (Recall that because \ characters are escaped, "/\\" refers to
 * /\, and "\\/" refers to \/.)
 * The 2x2 grid is as follows:
 * 
 * 
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * [
 * "//",
 * "/ "
 * ]
 * Output: 3
 * Explanation: The 2x2 grid is as follows:
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= grid.length == grid[0].length <= 30
 * grid[i][j] is either '/', '\', or ' '.
 * 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func regionsBySlashes(grid []string) int {
    
}
// @lc code=end
