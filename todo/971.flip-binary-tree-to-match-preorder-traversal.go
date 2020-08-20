package main


/*
 * @lc app=leetcode id=971 lang=golang
 *
 * [971] Flip Binary Tree To Match Preorder Traversal
 *
 * https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/description/
 *
 * algorithms
 * Medium (45.67%)
 * Likes:    264
 * Dislikes: 128
 * Total Accepted:    12.9K
 * Total Submissions: 28.2K
 * Testcase Example:  '[1,2]\n[2,1]'
 *
 * Given a binary tree with N nodes, each node has a different value from {1,
 * ..., N}.
 * 
 * A node in this binary tree can be flipped by swapping the left child and the
 * right child of that node.
 * 
 * Consider the sequence of N values reported by a preorder traversal starting
 * from the root.  Call such a sequence of N values the voyage of the tree.
 * 
 * (Recall that a preorder traversal of a node means we report the current
 * node's value, then preorder-traverse the left child, then preorder-traverse
 * the right child.)
 * 
 * Our goal is to flip the least number of nodes in the tree so that the voyage
 * of the tree matches the voyage we are given.
 * 
 * If we can do so, then return a list of the values of all nodes flipped.  You
 * may return the answer in any order.
 * 
 * If we cannot do so, then return the list [-1].
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: root = [1,2], voyage = [2,1]
 * Output: [-1]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * 
 * 
 * Input: root = [1,2,3], voyage = [1,3,2]
 * Output: [1]
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * 
 * 
 * Input: root = [1,2,3], voyage = [1,2,3]
 * Output: []
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= N <= 100
 * 
 * 
 * 
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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    
}
// @lc code=end
