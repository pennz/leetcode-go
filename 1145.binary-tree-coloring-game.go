package main

/*
 * @lc app=leetcode id=1145 lang=golang
 *
 * [1145] Binary Tree Coloring Game
 *
 * https://leetcode.com/problems/binary-tree-coloring-game/description/
 *
 * algorithms
 * Medium (50.83%)
 * Likes:    324
 * Dislikes: 64
 * Total Accepted:    14.9K
 * Total Submissions: 29.3K
 * Testcase Example:  '[1,2,3,4,5,6,7,8,9,10,11]\n11\n3'
 *
 * Two players play a turn based game on a binary tree.  We are given the root
 * of this binary tree, and the number of nodes n in the tree.  n is odd, and
 * each node has a distinct value from 1 to n.
 *
 * Initially, the first player names a value x with 1 <= x <= n, and the second
 * player names a value y with 1 <= y <= n and y != x.  The first player colors
 * the node with value x red, and the second player colors the node with value
 * y blue.
 *
 * Then, the players take turns starting with the first player.  In each turn,
 * that player chooses a node of their color (red if player 1, blue if player
 * 2) and colors an uncolored neighbor of the chosen node (either the left
 * child, right child, or parent of the chosen node.)
 *
 * If (and only if) a player cannot choose such a node in this way, they must
 * pass their turn.  If both players pass their turn, the game ends, and the
 * winner is the player that colored more nodes.
 *
 * You are the second player.  If it is possible to choose such a y to ensure
 * you win the game, return true.  If it is not possible, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
 * Output: true
 * Explanation: The second player can choose the node with value 2.
 *
 *
 *
 * Constraints:
 *
 *
 * root is the root of a binary tree with n nodes and distinct node values from
 * 1 to n.
 * n is odd.
 * 1 <= x <= n <= 100
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

/* Solution found online
https://csnotes.dev/en/post/leetcode/leetcode-1145/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	count := []int{0, 0}
	helper_func(root, x, count)
	//return btreeGameWinningMoveBruteForce(root, n, x) // with divide and conquer?
	return Max(Max(count[0], count[1]), n-count[0]-count[1]-1) > n/2
}

func btreeGameWinningMoveBruteForce(root *TreeNode, n int, x int) bool {
	return true
}

// some comments
func helper_func(root *TreeNode, x int, count []int) int {
	// test another
	// test another
	if root == nil {
		return 0
	}
	l := helper_func(root.Left, x, count)
	r := helper_func(root.Right, x, count)
	if root.Val == x {
		count[0] = l
		count[1] = r
	}
	return l + r + 1
}

// @lc code=end
