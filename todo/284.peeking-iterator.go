package main


/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 *
 * https://leetcode.com/problems/peeking-iterator/description/
 *
 * algorithms
 * Medium (45.71%)
 * Likes:    516
 * Dislikes: 378
 * Total Accepted:    101.1K
 * Total Submissions: 221.2K
 * Testcase Example:  '["PeekingIterator","next","peek","next","next","hasNext"]\n[[[1,2,3]],[],[],[],[],[]]'
 *
 * Given an Iterator class interface with methods: next() and hasNext(), design
 * and implement a PeekingIterator that support the peek() operation -- it
 * essentially peek() at the element that will be returned by the next call to
 * next().
 * 
 * Example:
 * 
 * 
 * Assume that the iterator is initialized to the beginning of the list:
 * [1,2,3].
 * 
 * Call next() gets you 1, the first element in the list.
 * Now you call peek() and it returns 2, the next element. Calling next() after
 * that still return 2. 
 * You call next() the final time and it returns 3, the last element. 
 * Calling hasNext() after that should return false.
 * 
 * 
 * Follow up: How would you extend your design to be generic and work with all
 * types, not just integer?
 * 
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    
}

func Constructor(iter *Iterator) *PeekingIterator {
    
}

func (this *PeekingIterator) hasNext() bool {
    
}

func (this *PeekingIterator) next() int {
    
}

func (this *PeekingIterator) peek() int {
    
}
// @lc code=end
