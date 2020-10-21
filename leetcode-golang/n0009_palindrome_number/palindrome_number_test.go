// go test -v .
package n0009_palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	if ans := isPalindrome(123); ans == true {
		t.Error("Failure: ", 123)
	}

	if ans := isPalindrome(0); ans == false {
		t.Error("Failure: ", 0)
	}

	if ans := isPalindrome(121); ans == false {
		t.Error("Failure: ", 121)
	}

	if ans := isPalindrome(-121); ans == true {
		t.Error("Failure: ", -121)
	}
}
