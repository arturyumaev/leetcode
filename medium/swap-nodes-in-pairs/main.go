package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// runtime: 1ms beats 78.10%
// memory: 2.19MB beats 99.81%
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := head.Next
	a := head
	for {
		b := a.Next
		c := b.Next
		b.Next = a
		a.Next = c
		temp := a
		a = c

		if a == nil {
			break
		}
		if a.Next == nil {
			break
		}

		temp.Next = a.Next
	}

	return newHead
}

func main() {
	swapPairs(nil)
}
