package main

/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 *
 * https://leetcode.com/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (84.88%)
 * Likes:    1879
 * Dislikes: 335
 * Total Accepted:    409.9K
 * Total Submissions: 482.9K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 * You're given strings J representing the types of stones that are jewels, and
 * S representing the stones you have.  Each character in S is a type of stone
 * you have.  You want to know how many of the stones you have are also
 * jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are
 * letters. Letters are case sensitive, so "a" is considered a different type
 * of stone from "A".
 *
 * Example 1:
 *
 *
 * Input: J = "aA", S = "aAAbbbb"
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: J = "z", S = "ZZ"
 * Output: 0
 *
 *
 * Note:
 *
 *
 * S and J will consist of letters and have length at most 50.
 * The characters in J are distinct.
 *
 *
 */

// @lc code=start

import "strings"

func numJewelsInStones(J string, S string) int {
	var cnt int = 0
	jSet := make(map[rune]bool, len(J))
	for _, j := range J {
		jSet[j] = true
	}
	for _, s := range S {
		if _, ok := jSet[s]; ok {
			cnt++
		}
	}
	return cnt
}
func numJewelsInStones2(J string, S string) int {
	var cnt int = 0
	for _, j := range J {
		cnt += strings.Count(S, string(j))
	}
	return cnt
}

// @lc code=end
