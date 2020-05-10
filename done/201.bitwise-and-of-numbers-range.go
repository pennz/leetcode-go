package main

//import "runtime"

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

/*
Learn from others: https://leetcode.com/problems/bitwise-and-of-numbers-range/discuss/56729/Bit-operation-solution(JAVA)

The idea is very simple:
1. last bit of (odd number & even number) is 0.
2. when m != n, There is at least an odd number and an even number, so the last bit position result is 0.
3. Move m and n rigth a position.

Keep doing step 1,2,3 until m equal to n, use a factor to record the iteration time.

public class Solution {
    public int rangeBitwiseAnd(int m, int n) {
        if(m == 0){
            return 0;
        }
        int moveFactor = 1;
		// eg.  1010, 1011 -> then we have it.
        while(m != n){
            m >>= 1;
            n >>= 1;
            moveFactor <<= 1;
        }
        return m * moveFactor;
    }
}
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
	q -= p // seems it is not reacheable, keep it here just in case
	return q
}

func getHighestBitValueMask(m int) int { // assume m >= 0
	mask := -1
	for i := 1; i <= m; i = i << 1 {
		mask = mask << 1
	}
	return mask
}
func getHighestBitValue(m int) int { // assume m >= 0
	// eg  0000 1001 -> 0000 1000
	// 1001 -> 11110110 + 1 (for -)
	// then &, we got 0000 0001
	p := 1 // find other zero ones
	for p <= m {
		p = p << 1
	}
	return p >> 1
}

func rangeBitwiseAnd(m int, n int) int {
	// just check for bits, if one is zero, can ignore others
	// just for bit representation, from to all will be zero
	if m == 0 {
		return 0
	}

	moves := 0
	for m != n {
		m >>= 1
		n >>= 1
		moves += 1
	}
	return m << moves
	// my solution is below, not elegant as the above one
	//if m == n {
	//	return m
	//}
	//// runtime.Breakpoint()
	//andV := n & m

	//hbv := andV & getHighestBitValueMask(n-m)

	//// return getHighestBitsValue(n)
	//return hbv
}

func rangeBitwiseAndForce(m int, n int) int {
	var andValue int = m
	for i := m + 1; i <= n; i++ {
		andValue &= i
	}
	return andValue
}

// @lc code=end
