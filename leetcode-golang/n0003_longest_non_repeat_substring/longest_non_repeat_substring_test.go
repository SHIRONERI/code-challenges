// go test -v -bench=.
package n0003_longest_non_repeat_substring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	maxLen := lengthOfLongestSubstring("abcabcbb")
	if maxLen != 3 {
		t.Error("Failure: ", maxLen)
	}

	maxLen = lengthOfLongestSubstring("bbbbb")
	if maxLen != 1 {
		t.Error("Failure: ", maxLen)
	}

	maxLen = lengthOfLongestSubstring("pwwkew")
	if maxLen != 3 {
		t.Error("Failure: ", maxLen)
	}

	maxLen = lengthOfLongestSubstring("abccba")
	if maxLen != 3 {
		t.Error("Failure: ", maxLen)
	}

	maxLen = lengthOfLongestSubstring("aab")
	if maxLen != 2 {
		t.Error("Failure: ", maxLen)
	}

	maxLen = lengthOfLongestSubstring("abcbabadefg")
	if maxLen != 6 {
		t.Error("Failure: ", maxLen)
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcdefghiailzjscvmcziinasddefgsdhiaqwttiilzjsfbwudjfhal")
	}
}
