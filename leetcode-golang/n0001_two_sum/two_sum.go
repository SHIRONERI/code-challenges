//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

package n0001_two_sum

// O(n/2) Solution
func twoSum(nums []int, target int) []int {
	find := make(chan []int, 1) // only one solution
	stop := make(chan struct{})

	// forward scanning
	go func() {
		numsLen := len(nums)
		expected := make(map[int]int, numsLen)
		for i := 0; i < numsLen; i++ {
			select {
			case <-stop:
				return
			default:
				currentNum := nums[i]
				if j, ok := expected[currentNum]; ok {
					// j must be smaller than i
					find <- []int{j, i}
					return
				} else {
					expected[target-currentNum] = i
				}
			}
		}
	}()

	// reverse scanning
	go func() {
		numsLen := len(nums)
		expected := make(map[int]int, numsLen)
		for i := numsLen - 1; i > -1; i-- {
			select {
			case <-stop:
				return
			default:
				currentNum := nums[i]
				if j, ok := expected[currentNum]; ok {
					// j must be greater than i
					find <- []int{i, j}
					return
				} else {
					expected[target-currentNum] = i
				}
			}
		}
	}()

	result := <-find
	close(stop)

	return result
}
