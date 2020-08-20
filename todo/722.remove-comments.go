package main

/*
 * @lc app=leetcode id=722 lang=golang
 *
 * [722] Remove Comments
 *
 * https://leetcode.com/problems/remove-comments/description/
 *
 * algorithms
 * Medium (34.59%)
 * Likes:    310
 * Dislikes: 906
 * Total Accepted:    28.5K
 * Total Submissions: 82.4K
 * Testcase Example:
 *
 * Given a C++ program, remove comments from it. The program source is an array
 * where source[i] is the i-th line of the source code.  This represents the
 * result of splitting the original source code string by the newline character
 * \n.
 *
 * In C++, there are two types of comments, line comments, and block comments.
 *
 * The string // denotes a line comment, which represents that it and rest of
 * the characters to the right of it in the same line should be ignored.
 *
 * The string / denotes a block comment, which represents that all characters
 * until the next (non-overlapping) occurrence of / should be ignored.  (Here,
 * occurrences happen in reading order: line by line from left to right.)  To
 * be clear, the string // does not yet end the block comment, as the ending
 * would be overlapping the beginning.
 *
 * The first effective comment takes precedence over others: if the string //
 * occurs in a block comment, it is ignored. Similarly, if the string /* occurs
 * in a line or block comment, it is also ignored.
 *
 * If a certain line of code is empty after removing comments, you must not
 * output that line: each string in the answer list will be non-empty.
 *
 * There will be no control characters, single quote, or double quote
 *
 * The line by line code is visualized as below:
 * int main()
 * {
 * ⁠ // variable declaration
 * int a, b, c;
 * ⁠  multiline
 * ⁠  comment for
 * a = b + c;
 * }
 *
 * Output: ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]
 *
 * The line by line code is visualized as below:
 * int main()
 * {
 * ⁠
 * int a, b, c;
 * a = b + c;
 * }
 *
 * Explanation:
 * The string /* denotes a block comment, including line 1 and lines 6-9. The
 * string // denotes line 4 as comments.
 *
 *
 *
 * Example 2:
 *
 * Input:
 *
 */

// @lc code=start
func removeComments(source []string) []string {

}

// @lc code=end
