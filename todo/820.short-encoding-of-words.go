package main
/*
 * @lc app=leetcode id=820 lang=golang
 *
 * [820] Short Encoding of Words
 *
 * https://leetcode.com/problems/short-encoding-of-words/description/
 *
 * algorithms
 * Medium (50.70%)
 * Likes:    251
 * Dislikes: 65
 * Total Accepted:    13.8K
 * Total Submissions: 27.3K
 * Testcase Example:  '["time", "me", "bell"]'
 *
 * Given a list of words, we may encode it by writing a reference string S and
 * a list of indexes A.
 * 
 * For example, if the list of words is ["time", "me", "bell"], we can write it
 * as S = "time#bell#" and indexes = [0, 2, 5].
 * 
 * Then for each index, we will recover the word by reading from the reference
 * string from that index until we reach a "#" character.
 * 
 * What is the length of the shortest reference string S possible that encodes
 * the given words?
 * 
 * Example:
 * 
 * 
 * Input: words = ["time", "me", "bell"]
 * Output: 10
 * Explanation: S = "time#bell#" and indexes = [0, 2, 5].
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= words.length <= 2000.
 * 1 <= words[i].length <= 7.
 * Each word has only lowercase letters.
 * 
 * 
 */

// @lc code=start
func minimumLengthEncoding(words []string) int {
    
}
// @lc code=end
