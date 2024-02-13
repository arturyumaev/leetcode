package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	emptyNode = &ListNode{}
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultHead := result

	n1 := l1
	n2 := l2

	overflow := 0

	for {
		v1 := n1.Val
		v2 := n2.Val

		sum := v1 + v2 + overflow
		overflow = 0
		if sum >= 10 {
			overflow = 1
			sum -= 10
		}

		result.Val = sum

		if n1.Next == nil && n2.Next == nil && overflow == 0 {
			break
		}

		n1 = n1.Next
		n2 = n2.Next

		if n1 == nil {
			n1 = emptyNode
		}
		if n2 == nil {
			n2 = emptyNode
		}

		result.Next = &ListNode{}
		result = result.Next
	}

	return resultHead
}
