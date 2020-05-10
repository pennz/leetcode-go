package main

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 * Input: [5,7]
 * Output: 4
 *
 * Example 2:
 *
 * Input: [0,1]
 * Output: 0
 */

// @lc code=start
// getHighestBitsValue it just have no good meaning... I will revert back my previous solution
func getHighestBitsValue(m int) int { // assume m >= 0
	// eg  0000 1001 -> 0000 1000
	// 1001 -> 11110110 + 1 (for -)
	// then &, we got 0000 0001
	p := 1 // find other zero ones
	for p <= m {
		p = p << 1
	}
	p = p >> 1
	q := p
	for q <= m {
		p = p >> 1
		if p == 0 {
			return q
		}
		if p&m == 0 {
			return q
		}
		q += p
	}
	q -= p
	return q
}

func rangeBitwiseAnd(m int, n int) int {
	// just check for bits, if one is zero, can ignore others
	// just for bit representation, from to all will be zero
	if m == n {
		return m
	}
	hbv := getHighestBitsValue(n & m)

	if m < hbv {
		return 0
	}
	// return getHighestBitsValue(n)
	return hbv
}

func rangeBitwiseAndForce(m int, n int) int {
	var andValue int = m
	for i := m + 1; i <= n; i++ {
		andValue &= i
	}
	return andValue
}

// @lc code=end
