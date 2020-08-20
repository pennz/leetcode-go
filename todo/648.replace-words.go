package main
/*
 * @lc app=leetcode id=648 lang=golang
 *
 * [648] Replace Words
 *
 * https://leetcode.com/problems/replace-words/description/
 *
 * algorithms
 * Medium (56.56%)
 * Likes:    751
 * Dislikes: 131
 * Total Accepted:    56.5K
 * Total Submissions: 99.9K
 * Testcase Example:  '["cat","bat","rat"]\n"the cattle was rattled by the battery"'
 *
 * In English, we have a concept called root, which can be followed by some
 * other words to form another longer word - let's call this word successor.
 * For example, the root an, followed by other, which can form another word
 * another.
 * 
 * Now, given a dictionary consisting of many roots and a sentence. You need to
 * replace all the successor in the sentence with the root forming it. If a
 * successor has many roots can form it, replace it with the root with the
 * shortest length.
 * 
 * You need to output the sentence after the replacement.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: dict = ["cat","bat","rat"], sentence = "the cattle was rattled by the
 * battery"
 * Output: "the cat was rat by the bat"
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The input will only have lower-case letters.
 * 1 <= dict.length <= 1000
 * 1 <= dict[i].length <= 100
 * 1 <= sentence words number <= 1000
 * 1 <= sentence words length <= 1000
 * 
 * 
 */

// @lc code=start
func replaceWords(dict []string, sentence string) string {
    
}
// @lc code=end
