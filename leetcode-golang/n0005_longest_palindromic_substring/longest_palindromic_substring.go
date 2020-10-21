//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example 1:
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//
//Example 2:
//
//Input: "cbbd"
//Output: "bb"
package n0005_longest_palindromic_substring

const stringMaxLen = 1000

// O(n^2) Time & O(n^2) Space DP Solution
func longestPalindrome(s string) string {
	stringLen, maxPalindromeLen, palindromeBegin := len(s), 1, 0
	if stringLen == 0 {
		return ""
	}

	// dp[i, j] indicate the s[i : j] is palindrome string or not
	// dp[i, j] = true                                if i == j
	//          = s[i] == s[j]                        if j = i + 1
	//          = s[i] == s[j] && dp[i + 1][j - 1]    if j > i + 1
	var dp [stringMaxLen][stringMaxLen]bool
	for i := 0; i < stringLen; i++ {
		dp[i][i] = true
	}

	for i := 0; i < stringLen-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			palindromeBegin = i
			maxPalindromeLen = 2
		}
	}

	for pLen := 3; pLen <= stringLen; pLen++ {
		for i := 0; i+pLen-1 < stringLen; i++ {
			j := i + pLen - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				palindromeBegin = i
				maxPalindromeLen = pLen
			}
		}
	}

	return s[palindromeBegin : palindromeBegin+maxPalindromeLen]
}

// O(n) Solution based on Manacher's Algorithm
// https://articles.leetcode.com/longest-palindromic-substring-part-ii/
//func longestPalindrome(s string) string {
//
//}
