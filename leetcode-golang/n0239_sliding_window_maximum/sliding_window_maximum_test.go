package n0239_sliding_window_maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	var ret []int
	ret = maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, ret)
	ret = maxSlidingWindow([]int{0, 1, 2, 3, 5, 3, 2, 1}, 3)
	assert.Equal(t, []int{2, 3, 5, 5, 5, 3}, ret)
	ret = maxSlidingWindow([]int{1}, 1)
	assert.Equal(t, []int{1}, ret)
	ret = maxSlidingWindow([]int{1, -1}, 2)
	assert.Equal(t, []int{1}, ret)
	ret = maxSlidingWindow([]int{1, -1}, 1)
	assert.Equal(t, []int{1, -1}, ret)
	ret = maxSlidingWindow([]int{7, 2, 4}, 2)
	assert.Equal(t, []int{7, 4}, ret)
	ret = maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, ret)
}
