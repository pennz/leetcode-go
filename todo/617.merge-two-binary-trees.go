package main

/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
 *
 * https://leetcode.com/problems/merge-two-binary-trees/description/
 *
 * algorithms
 * Easy (74.13%)
 * Likes:    3186
 * Dislikes: 175
 * Total Accepted:    308.9K
 * Total Submissions: 416.7K
 * Testcase Example:  '[1,3,2,5]\n[2,1,3,null,4,null,7]'
 *
 * Given two binary trees and imagine that when you put one of them to cover
 * the other, some nodes of the two trees are overlapped while the others are
 * not.
 * 
 * You need to merge them into a new binary tree. The merge rule is that if two
 * nodes overlap, then sum node values up as the new value of the merged node.
 * Otherwise, the NOT null node will be used as the node of new tree.
 * 
 * Example 1:
 * 
 * 
 * Input: 
 * Tree 1                     Tree 2                  
 * ⁠         1                         2                             
 * ⁠        / \                       / \                            
 * ⁠       3   2                     1   3                        
 * ⁠      /                           \   \                      
 * ⁠     5                             4   7                  
 * Output: 
 * Merged tree:
 * 3
 * / \
 * 4   5
 * / \   \ 
 * 5   4   7
 * 
 * 
 * 
 * 
 * Note: The merging process must start from the root nodes of both trees.
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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    
}
// @lc code=end
