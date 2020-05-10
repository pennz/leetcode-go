package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (27.81%)
 * Likes:    1619
 * Dislikes: 194
 * Total Accepted:    167.7K
 * Total Submissions: 602.8K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non negative stringegers, arrange them such that they form the
 * largest number.
 *
 * Example 1:
 *
 *
 * Input: [10,2]
 * Output: "210"
 *
 * Example 2:
 *
 *
 * Input: [3,30,34,5,9]
 * Output: "9534330"
 *
 *
 * Note: The result may be very large, so you need to return a string instead
 * of an stringeger.
 *
 */

// @lc code=start
func quickSort(nums []string, compareFunc func(string, string) bool) {
	if len(nums) <= 1 {
		return
	}

	pivot := partition(nums, compareFunc)
	ls, rs := nums[:pivot], nums[pivot:]
	quickSort(ls, compareFunc)
	quickSort(rs, compareFunc)
}
func partition(nums []string, compareFunc func(string, string) bool) int {
	pivot, pv := 0, nums[0]
	if pv == "" {
		pivot += 1
	}
	return pivot
}

// compareFunc if n1s > n2s, return true, we can just assume the padding zeros
// e.g. 53 530, true; 530 53 -> false ;  52, 530, false; 520, 53 -> false
func compareFunc(n1s string, n2s string) bool {
	if n1s == n2s { // for corner case "0" "0"
		return false
	}
	n1l, n2l := len(n1s), len(n2s)
	minl := n1l
	longer1 := true
	if n1l > n2l {
		minl = n2l
	} else {
		longer1 = false
	}
	pivot := 0

	for pivot = 0; pivot < minl; pivot++ { // oh... i just use pivot:= 0, local variable shadows previous one
		if n1s[pivot] > n2s[pivot] {
			return true
		} else if n1s[pivot] == n2s[pivot] {
			continue
		} else {
			return false
		}
	}
	// jus compare from the pivot place
	// only reach here if it is equal for range[:minl]
	// so 530 53, but 531 53 -> true, 530 53 -> false
	// 5301 53 , (so this is not a good way?)
	// 121 and 12, we compare and found 12 the same, then we compare [12] 121 -> with [12]1 12, and check which one is better
	// so we should compare 121 and leftover then our number, then we print out which one is what we want
	if longer1 {
		return compareFunc(n1s[pivot:]+n2s, n1s) // 121 and 12, [12]1 12 with [12] 121
	}
	return compareFunc(n2s, n2s[pivot:]+n1s) // eg: 12 and 121, then we compare [12]121 and [12]1 12

}

func largestNumber(nums []int) string {
	// just sort by number, and func is num
	var res string
	var numss = make([]string, len(nums))
	for i, n := range nums {
		numss[i] = strconv.Itoa(n)
	}
	sort.Slice(numss, func(i, j int) bool {
		return compareFunc(numss[i], numss[j])
	})
	for _, s := range numss {
		res += s
	}
	if res[0:1] == "0" {
		return "0"
	}
	return res
}

// @lc code=end
