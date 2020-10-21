package n0206_reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n1.Next = n2
	n2.Next = n3
	printListNode(n1)
	printListNode(reverseList(n1))
}

func printListNode(node *ListNode) {
	for {
		fmt.Println(node.Val)
		if node.Next != nil {
			node = node.Next
		} else {
			return
		}
	}
}
