// go test -v .
package n0002_add_two_numbers

import (
	"strconv"
	"testing"
)

func BigNumberToString(node *ListNode) string {
	cursor, bignumber := node, ""
	for cursor.Next != nil {
		bignumber += strconv.Itoa(cursor.Val)
		cursor = cursor.Next
	}
	return bignumber
}

func TestAddTwoNumbers1(t *testing.T) {
	// (4, 3, 2, 1) + (9, 8, 7) = (3, 2, 0, 2)

	l1_1, l1_2, l1_3, l1_4 :=
		&ListNode{4, nil}, &ListNode{3, nil}, &ListNode{2, nil}, &ListNode{1, nil}
	l1_1.Next = l1_2
	l1_2.Next = l1_3
	l1_3.Next = l1_4

	l2_1, l2_2, l2_3 :=
		&ListNode{9, nil}, &ListNode{8, nil}, &ListNode{7, nil}
	l2_1.Next = l2_2
	l2_2.Next = l2_3

	result := addTwoNumbers(l1_1, l2_1)

	if result.Val != 3 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Val != 2 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Next.Val != 0 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Next.Next.Val != 2 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Next.Next.Next != nil {
		t.Error("Failure: ", BigNumberToString(result))
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	// (9, 9) + (1) = (0, 0, 1)

	l1_1, l1_2 := &ListNode{9, nil}, &ListNode{9, nil}
	l1_1.Next = l1_2

	l2_1 := &ListNode{1, nil}

	result := addTwoNumbers(l1_1, l2_1)

	if result.Val != 0 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Val != 0 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Next.Val != 1 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next.Next.Next != nil {
		t.Error("Failure: ", BigNumberToString(result))
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	// () + (1) = (1)

	l2_1 := &ListNode{1, nil}

	result := addTwoNumbers(nil, l2_1)

	if result.Val != 1 {
		t.Error("Failure: ", BigNumberToString(result))
	}

	if result.Next != nil {
		t.Error("Failure: ", BigNumberToString(result))
	}
}
