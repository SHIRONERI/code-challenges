//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//
//Example 2:
//
//Input: -123
//Output: -321
//
//Example 3:
//
//Input: 120
//Output: 21
//
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
//For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
package n0007_reverse_integer

// int32: [-2147483648, 2147483647]
const Int32Min = -1 << 31
const Int32Max = 1<<31 - 1

func reverse(x int) int {
	var reverseInteger int
	for x != 0 {
		bit := x % 10
		x /= 10
		if reverseInteger > Int32Max/10 || (reverseInteger == Int32Min/10 && bit > 7) {
			return 0
		}
		if reverseInteger < Int32Min/10 || (reverseInteger == Int32Min/10 && bit < -8) {
			return 0
		}
		reverseInteger = reverseInteger*10 + bit
	}
	return reverseInteger
}
