// go test -v -bench=.
package n0001_two_sum

import (
	"reflect"
	"testing"
)

var nums = []int{
	1, 517, 522, 18, 540, 29, 30, 39, 557, 575, 583, 585, 589, 953, 609, 100, 620, 625, 632, 126, 641, 643, 139, 156,
	669, 159, 672, 880, 177, 182, 184, 200, 716, 206, 208, 722, 211, 212, 729, 732, 223, 224, 227, 740, 229, 759, 250,
	254, 258, 265, 269, 283, 292, 297, 812, 305, 830, 320, 837, 841, 330, 334, 335, 860, 355, 368, 370, 381, 385, 396,
	399, 401, 404, 922, 413, 420, 948, 952, 441, 959, 961, 966, 971, 974, 975, 475, 476, 991, 483, 486, 498,
}
var target = 1500
var result = []int{16, 27}

func TestTwoSum(t *testing.T) {
	ret := twoSum(nums, target)
	if !reflect.DeepEqual(ret, result) {
		t.Error("TwoSum Failure")
	}
}

func BenchmarkTwoSum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
