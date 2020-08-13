/*
 * @lc app=leetcode id=903 lang=golang
 *
 * [903] Valid Permutations for DI Sequence
 *
 * https://leetcode.com/problems/valid-permutations-for-di-sequence/description/
 *
 * algorithms
 * Hard (49.72%)
 * Likes:    288
 * Dislikes: 28
 * Total Accepted:    5.9K
 * Total Submissions: 11.8K
 * Testcase Example:  '"DID"'
 *
 * We are given S, a length n string of characters from the set {'D', 'I'}.
 * (These letters stand for "decreasing" and "increasing".)
 * 
 * A valid permutation is a permutation P[0], P[1], ..., P[n] of integers {0,
 * 1, ..., n}, such that for all i:
 * 
 * 
 * If S[i] == 'D', then P[i] > P[i+1], and;
 * If S[i] == 'I', then P[i] < P[i+1].
 * 
 * 
 * How many valid permutations are there?  Since the answer may be large,
 * return your answer modulo 10^9 + 7.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "DID"
 * Output: 5
 * Explanation: 
 * The 5 valid permutations of (0, 1, 2, 3) are:
 * (1, 0, 3, 2)
 * (2, 0, 3, 1)
 * (2, 1, 3, 0)
 * (3, 0, 2, 1)
 * (3, 1, 2, 0)
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= S.length <= 200
 * S consists only of characters from the set {'D', 'I'}.
 * 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func numPermsDISequence(S string) int {
    
}
// @lc code=end
