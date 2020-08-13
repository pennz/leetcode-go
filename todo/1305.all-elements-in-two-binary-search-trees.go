/*
 * @lc app=leetcode id=1305 lang=golang
 *
 * [1305] All Elements in Two Binary Search Trees
 *
 * https://leetcode.com/problems/all-elements-in-two-binary-search-trees/description/
 *
 * algorithms
 * Medium (76.15%)
 * Likes:    387
 * Dislikes: 19
 * Total Accepted:    31.7K
 * Total Submissions: 41.7K
 * Testcase Example:  '[2,1,4]\r\n[1,0,3]\r'
 *
 * Given two binary search trees root1 and root2.
 * 
 * Return a list containing all the integers from both trees sorted in
 * ascending order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root1 = [2,1,4], root2 = [1,0,3]
 * Output: [0,1,1,2,3,4]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
 * Output: [-10,0,0,1,2,5,7,10]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: root1 = [], root2 = [5,1,7,0,2]
 * Output: [0,1,2,5,7]
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: root1 = [0,-10,10], root2 = []
 * Output: [-10,0,10]
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: root1 = [1,null,8], root2 = [8,1]
 * Output: [1,1,8,8]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * Each tree has at most 5000 nodes.
 * Each node's value is between [-10^5, 10^5].
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    
}
// @lc code=end
