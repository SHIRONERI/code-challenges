// go test -v -bench=.
package n0004_median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	if ans := findMedianSortedArrays([]int{1, 2, 3, 5, 6}, []int{4}); ans != 3.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{4}, []int{1, 2, 3, 5, 6}); ans != 3.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2}); ans != 4.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{2}, []int{1, 3, 5, 7, 9}); ans != 4.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 6}); ans != 5.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{2, 6}, []int{1, 3, 5, 7, 9}); ans != 5.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); ans != 2.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{3, 4}, []int{1, 2}); ans != 2.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1}, []int{3, 4}); ans != 3.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{3, 4}, []int{1}); ans != 3.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1, 2}, []int{}); ans != 1.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{}, []int{1, 2}); ans != 1.5 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{1, 2, 3}, []int{}); ans != 2.0 {
		t.Error("Failure: ", ans)
	}

	if ans := findMedianSortedArrays([]int{}, []int{1, 2, 3}); ans != 2.0 {
		t.Error("Failure: ", ans)
	}
}
