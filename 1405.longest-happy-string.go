package main

import "strings"

/*
 * @lc app=leetcode id=1405 lang=golang
 *
 * [1405] Longest Happy String
 *
 * https://leetcode.com/problems/longest-happy-string/description/
 *
 * algorithms
 * Medium (45.50%)
 * Likes:    155
 * Dislikes: 56
 * Total Accepted:    6.5K
 * Total Submissions: 14.2K
 * Testcase Example:  '1\n1\n7'
 *
 * A string is called happy if it does not have any of the strings 'aaa', 'bbb'
 * or 'ccc' as a substring.
 *
 * Given three integers a, b and c, return any string s, which satisfies
 * following conditions:
 *
 *
 * s is happy and longest possible.
 * s contains at most a occurrences of the letter 'a', at most b occurrences of
 * the letter 'b' and at most c occurrences of the letter 'c'.
 * s will only contain 'a', 'b' and 'c' letters.
 *
 *
 * If there is no such string s return the empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: a = 1, b = 1, c = 7
 * Output: "ccaccbcc"
 * Explanation: "ccbccacc" would also be a correct answer.
 *
 *
 * Example 2:
 *
 *
 * Input: a = 2, b = 2, c = 1
 * Output: "aabbc"
 *
 *
 * Example 3:
 *
 *
 * Input: a = 7, b = 1, c = 0
 * Output: "aabaa"
 * Explanation: It's the only correct answer in this case.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= a, b, c <= 100
 * a + b + c > 0
 *
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=start
// longestDiverseString use greedy alg, put the way to max len (generation way)
func putString(lenMap map[string]int) string {
	var ans strings.Builder

	for lenMap["a"] >= 0 || lenMap["b"] >= 0 || lenMap["c"] >= 0 {
		k := findLongest(lenMap)
		ans.WriteString(k)
		lenMap[k]--
	}
	return ans.String()
}

func findLongest(lenm map[string]int) string {
	initFlag := false
	var vv int
	var kk string
	for k, v := range lenm {
		if !initFlag {
			vv = v
			kk = k
		} else {
			if v > vv {
				kk = k
			}
		}
	}
	return kk
}

func longestDiverseString(a int, b int, c int) string {
	// https://zxi.mytechroad.com/blog/greedy/leetcode-1405-longest-happy-string/
	// Solution: Greedy (dynamic programming)
	//Put the char with highest frequency first if its consecutive length of that char is < 2
	//or put one char if any of other two chars has consecutive length of 2.
	//
	//increase the consecutive length of itself and reset that for other two chars.
	//
	//Time complexity: O(n)
	//Space complexity: O(1)
	lenMap := make(map[string]int)
	consec := make(map[string]int)
	lenMap["a"] = a
	lenMap["b"] = b
	lenMap["c"] = c
	consec["a"] = 0
	consec["b"] = 0
	consec["c"] = 0
	//pivot := ""
	// we can use recursive , solve this issue
	return putString(lenMap)
	// find the one with most one
	// put two of the most one
	// put one of the other color
}

// MyLongestDiverseString not very good solution, hard to patch and patch and patch
func MyLongestDiverseString(a int, b int, c int) string {
	// happy -> no aaa bbb ccc
	// longest -> atmost a "a" ...
	// analyze: how to put things to position, permutation ~
	// We divide:
	// - a, b, c, maxV > 2*remV + 2
	// if not, maxV

	var ans strings.Builder

	sum := a + b + c
	maxV := max(max(a, b), c)
	remV := sum - maxV

	if maxV > 2*remV+2 {
		ans.WriteString("aa")
		return ans.String()
	}
	return "failed"
}

// @lc code=end
