//Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
//Example 1:
//
//Input: 121, Output: true
//
//Example 2:
//
//Input: -121, Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
//Example 3:
//
//Input: 10, Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//Follow up:
//
//Could you solve it without converting the integer to a string?
//
package n0009_palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp, palindrome := x, 0
	for tmp != 0 {
		bit := tmp % 10
		tmp /= 10
		palindrome = 10*palindrome + bit
	}

	return palindrome == x
}
