//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//You may assume nums1 and nums2 cannot be both empty.
//
//Example 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//
//Example 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5

package n0004_median_of_two_sorted_arrays

// Hints:
//
// To solve this problem, we need to understand "What is the use of median".
// In statistics, the median is used for dividing a set into two equal length subsets,
// that one subset is always greater than the other.
// If we understand the use of median for dividing, we are very close to the answer.

// O(log(min(m, n))) Solution
//
// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2481/Share-my-O(log(min(mn))-solution-with-explanation
//
// nums1: A1, A2, A3, ..., Am
// nums2: B1, B2, B3, ..., Bn
//
// the key point is to find a i in [0, m] and a j in [0, n], which makes i + j = m + n + 1 / 2
// so we get two sets of numbers as follow,
// Left Part             |   Right Part
// A1, A2, A3, ..., Ai   |   Ai+1, Ai+2, Ai+3, ..., Am
// B1, B2, B3, ..., Bj   |   Bj+1, Bj+2, Bj+3, ..., Bn
// if Ai <= Bj+1 and Bj <= Ai+1,
// than we have two sets of numbers which any number in right part is greater than left part
//
// Left Part             |   Right Part
//                       |   Ai+1, Ai+2, Ai+3, ..., Am
// B1, B2, B3, ..., Bj   |   Bj+1, Bj+2, Bj+3, ..., Bn
// if i == 0 means A is the right part, in that case, j + 1 will be out of range
//
// Left Part             |   Right Part
// A1, A2, A3, ..., Ai   |
// B1, B2, B3, ..., Bj   |   Bj+1, Bj+2, Bj+3, ..., Bn
// if i == m means A is the left part, in that case, i + 1 will be out of range
//
// in that condition,
// * if m + n is even
//     the median is: max(Ai, Bj) + min(Ai+1, Bj+1) / 2
// * if m + n is odd
//     the median is: max(Ai, Bj)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2

	if m == 0 || n == 0 {
		nums := append(nums1, nums2...)
		if len(nums)%2 == 1 {
			return float64(nums[halfLen-1])
		} else {
			return (float64(nums[halfLen-1]) + float64(nums[halfLen])) / 2.0
		}
	}

	// make sure m < n, so j = halfLen - i is always greater than zero
	var A, B []int
	if m < n {
		A, B = nums1, nums2
	} else {
		m, n = n, m
		A, B = nums2, nums1
	}

	// find i in [0, m]
	// especial, if i == 0 means A is the right part, if i == m means A is the left part
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i

		if i > 0 && j >= 0 && j < n && A[i-1] > B[j] {
			// Ai > Bj+1, i need be smaller
			iMax = i - 1
		} else if j > 0 && i >= 0 && i < m && B[j-1] > A[i] {
			// Bj > Ai+1, i need be greater
			iMin = i + 1
		} else {
			var leftPartMax, rightPartMin float64
			if i == 0 {
				leftPartMax = float64(B[j-1])
			} else if j == 0 {
				leftPartMax = float64(A[i-1])
			} else {
				leftPartMax = float64(max(A[i-1], B[j-1]))
			}

			if (m+n)%2 == 1 {
				// m + n is odd
				return leftPartMax
			}

			if i == m {
				rightPartMin = float64(B[j])
			} else if j == n {
				rightPartMin = float64(A[i])
			} else {
				rightPartMin = float64(min(A[i], B[j]))
			}

			// m + n is even
			return (leftPartMax + rightPartMin) / 2.0
		}
	}

	return -1.0
}
