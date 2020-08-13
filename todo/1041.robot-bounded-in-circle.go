/*
 * @lc app=leetcode id=1041 lang=golang
 *
 * [1041] Robot Bounded In Circle
 *
 * https://leetcode.com/problems/robot-bounded-in-circle/description/
 *
 * algorithms
 * Medium (49.62%)
 * Likes:    280
 * Dislikes: 131
 * Total Accepted:    16.7K
 * Total Submissions: 33.7K
 * Testcase Example:  '"GGLLGG"'
 *
 * On an infinite plane, a robot initially stands at (0, 0) and faces north.
 * The robot can receive one of three instructions:
 * 
 * 
 * "G": go straight 1 unit;
 * "L": turn 90 degrees to the left;
 * "R": turn 90 degress to the right.
 * 
 * 
 * The robot performs the instructions given in order, and repeats them
 * forever.
 * 
 * Return true if and only if there exists a circle in the plane such that the
 * robot never leaves the circle.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "GGLLGG"
 * Output: true
 * Explanation: 
 * The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to
 * (0,0).
 * When repeating these instructions, the robot remains in the circle of radius
 * 2 centered at the origin.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "GG"
 * Output: false
 * Explanation: 
 * The robot moves north indefinitely.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "GL"
 * Output: true
 * Explanation: 
 * The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) ->
 * ...
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= instructions.length <= 100
 * instructions[i] is in {'G', 'L', 'R'}
 * 
 * 
 */

// @lc code=start
func isRobotBounded(instructions string) bool {
    
}
// @lc code=end
