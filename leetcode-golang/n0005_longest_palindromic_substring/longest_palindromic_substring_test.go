// go test -v .
package n0005_longest_palindromic_substring

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	if ans := longestPalindrome(""); ans != "" {
		t.Error("Failure: ", "", ans)
	}

	if ans := longestPalindrome("a"); ans != "a" {
		t.Error("Failure: ", "a", ans)
	}

	if ans := longestPalindrome("ccc"); ans != "ccc" {
		t.Error("Failure: ", "ccc", ans)
	}

	if ans := longestPalindrome("babad"); !(ans == "bab" || ans == "aba") {
		t.Error("Failure: ", "babad", ans)
	}

	if ans := longestPalindrome("cbbd"); !(ans == "bb") {
		t.Error("Failure: ", "cbbd", ans)
	}
}
