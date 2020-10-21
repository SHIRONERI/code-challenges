//Given a string, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package n0003_longest_non_repeat_substring

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n) time O(1) space Solution
func lengthOfLongestSubstring(s string) int {
	var chPosition [256]int // [0, 0, 0, ...]
	maxLength, substringLen, lastRepeatPos := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			// record current max substring length
			maxLength = max(substringLen, maxLength)

			// update characters position
			chPosition[s[i]] = i + 1

			// update last repeat character position
			lastRepeatPos = max(pos, lastRepeatPos)

			// update the current substring from last repeat character
			substringLen = i + 1 - lastRepeatPos
		} else {
			substringLen += 1
			chPosition[s[i]] = i + 1
		}
	}

	return max(maxLength, substringLen)
}
