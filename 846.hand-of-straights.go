package main

import "sort"
import "log"

/*
 * @lc app=leetcode id=846 lang=golang
 *
 * [846] Hand of Straights
 *
 * https://leetcode.com/problems/hand-of-straights/description/
 *
 * algorithms
 * Medium (54.20%)
 * Likes:    691
 * Dislikes: 82
 * Total Accepted:    48.6K
 * Total Submissions: 89.7K
 * Testcase Example:  '[1,2,3,6,2,3,4,7,8]\n3'
 *
 * Alice has a hand of cards, given as an array of integers.
 *
 * Now she wants to rearrange the cards into groups so that each group is size
 * W, and consists of W consecutive cards.
 *
 * Return true if and only if she can.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
 * Output: true
 * Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
 *
 * Example 2:
 *
 *
 * Input: hand = [1,2,3,4,5], W = 4
 * Output: false
 * Explanation: Alice's hand can't be rearranged into groups of 4.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= hand.length <= 10000
 * 0 <= hand[i] <= 10^9
 * 1 <= W <= hand.length
 *
 * Note: This question is the same as 1296:
 * https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
 */

// @lc code=start
func countMap(hand []int) (map[int]int, []int) {
	cm := make(map[int]int)
	for _, v := range hand {
		cm[v]++
	}
	keys := make([]int, 0, len(cm))
	for k := range cm {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return cm, keys
}

// isNStraightHand first get stat of number for each number of cards, then from
// lest one, count group until to the largest. Check if there is any card that
// do not fit in a group.
func isNStraightHand(hand []int, W int) bool {
	// simple thought: sort then check group
	// quicker way?
	if len(hand)%W != 0 {
		return false
	}
	cm, keys := countMap(hand)
	return checkHand(cm, W, keys)
}

// checkHand can just use for loop
func checkHand(cm map[int]int, W int, keys []int) bool {

	// key, just use the map to check if there are any left, no need this patch

	//toCheck := keys[0] // better to check before use
	//toRemove := cm[toCheck]
	//if len(keys) < W {
	//	for _, k := range keys {
	//		if cm[k] > 0 {
	//			return false
	//		}
	//	}
	//}
	keysToCheck := keys
	for len(keysToCheck) > 0 { // we just move along the way, so no need to handle the
		// last few number, it is just a patch
		checkNow := keysToCheck[0]

		delete(cm, checkNow)
		log.Println(keysToCheck)
		keysDoneCheck := 1

		log.Println(checkNow)

		checkNow++
		for i, toRemove := 1, cm[checkNow]; i < W; i, checkNow = i+1, checkNow+1 {
			if _, ok := cm[checkNow]; !ok {
				return false
			}
			cm[checkNow] -= toRemove
			switch remain := cm[checkNow]; {
			case remain < 0:
				return false
			case remain == 0:
				delete(cm, checkNow)
				keysDoneCheck++
			}
		}

		keysToCheck = keysToCheck[keysDoneCheck:]
	}
	return true
}

// myCheckHand is my recursive implementation, not very efficient
func myCheckHand(cm map[int]int, W int, keys []int) bool {
	if len(keys) < W {
		for _, k := range keys {
			if cm[k] > 0 {
				return false
			}
		}
		return true
	}
	toCheck := keys[0] // better to check before use
	toRemove := cm[toCheck]

	if toRemove != 0 {
		cm[toCheck] = 0 // remove from hand
		toCheck++
		for i := 1; i < W; i, toCheck = i+1, toCheck+1 {
			if _, ok := cm[toCheck]; !ok {
				return false
			}

			cm[toCheck] -= toRemove
			if cm[toCheck] < 0 {
				return false
			}
		}
	}
	return myCheckHand(cm, W, keys[1:])
}

// @lc code=end

/* Reference solution
Hand of Straights

Python Solution with Explanation

https://leetcode.com/problems/hand-of-straights/discuss/323274

* Lang:    python
* Author:  whissely
* Votes:   1

O(n * (n/w)) - see solution tab, the extra contribution is from searching the
minimum One improvement could be to go in sorted order
Space: O(n) - for the counter
Algorithm: First check if we have any remainders, if yes, return False. Build
a counter and while there are values in it, simply begin constructing groups,
at every iteration remove an occurence of that value, delete it if the element
count reaches 0, and look for the next consecutive number.
```
def isNStraightHand(self, hand, W):
        """
        :type hand: List[int]
        :type W: int
        :rtype: bool
        """
        if len(hand) % W:
            return False

        counter = collections.Counter(hand)
        while counter:
            elem = min(counter)
            cnt = 0
            while cnt < W:
                if elem not in counter:
                    return False
                counter[elem] -= 1
                if not counter[elem]:
                    del counter[elem]
                elem += 1
                cnt += 1

        return True
*/
