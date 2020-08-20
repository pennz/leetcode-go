package main


/*
 * @lc app=leetcode id=1349 lang=golang
 *
 * [1349] Maximum Students Taking Exam
 *
 * https://leetcode.com/problems/maximum-students-taking-exam/description/
 *
 * algorithms
 * Hard (41.96%)
 * Likes:    295
 * Dislikes: 9
 * Total Accepted:    5.6K
 * Total Submissions: 13.3K
 * Testcase Example:  '[["#",".","#","#",".","#"],[".","#","#","#","#","."],["#",".","#","#",".","#"]]'
 *
 * Given a m * n matrix seats  that represent seats distributions in a
 * classroom. If a seat is broken, it is denoted by '#' character otherwise it
 * is denoted by a '.' character.
 * 
 * Students can see the answers of those sitting next to the left, right, upper
 * left and upper right, but he cannot see the answers of the student sitting
 * directly in front or behind him. Return the maximum number of students that
 * can take the exam together without any cheating being possible..
 * 
 * Students must be placed in seats in good condition.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: seats = [["#",".","#","#",".","#"],
 * [".","#","#","#","#","."],
 * ["#",".","#","#",".","#"]]
 * Output: 4
 * Explanation: Teacher can place 4 students in available seats so they don't
 * cheat on the exam. 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: seats = [[".","#"],
 * ["#","#"],
 * ["#","."],
 * ["#","#"],
 * [".","#"]]
 * Output: 3
 * Explanation: Place all students in available seats. 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: seats = [["#",".",".",".","#"],
 * [".","#",".","#","."],
 * [".",".","#",".","."],
 * [".","#",".","#","."],
 * ["#",".",".",".","#"]]
 * Output: 10
 * Explanation: Place students in available seats in column 1, 3 and 5.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * seats contains only characters '.' and'#'.
 * m == seats.length
 * n == seats[i].length
 * 1 <= m <= 8
 * 1 <= n <= 8
 * 
 * 
 */

// @lc code=start
func maxStudents(seats [][]byte) int {
    
}
// @lc code=end
