//Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.
//
//Follow up:
//Could you solve it in linear time?
//
//Example:
//
//Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
//Output: [3,3,5,5,6,7]
//Explanation:
//
//Window position                Max
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
//
//
//Constraints:
//
//1 <= nums.length <= 10^5
//-10^4 <= nums[i] <= 10^4
//1 <= k <= nums.length
package n0239_sliding_window_maximum

// O(n)
func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{}
	}

	// Handling the first window (0, 1, k-1)
	deque, result := make([]int, 0, k), make([]int, 0, numsLen-k+1)
	for i := 0; i < k; i++ {
		deque = push(deque, nums, i)
	}
	result = append(result, nums[deque[0]])

	// Sliding the window (k, k+1, ..., len(nums)-1)
	for i := k; i < numsLen; i++ {
		// Check the front element of deque is out of window or not
		if i-deque[0] >= k {
			deque = deque[1:]
		}

		deque = push(deque, nums, i)

		// Fetch the maximum value
		result = append(result, nums[deque[0]])
	}

	return result
}

// push index of element into deque
func push(deque []int, nums []int, idx int) []int {
	// Keep the maximum in the front of deque
	for len(deque) > 0 {
		dequeLen := len(deque)
		backElement := deque[dequeLen-1]
		if nums[idx] > nums[backElement] {
			deque = deque[:dequeLen-1]
		} else {
			break
		}
	}
	return append(deque, idx)
}
