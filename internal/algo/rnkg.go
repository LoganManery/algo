package algo

// Reverse Nodes in K Group

type Node struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Edge cases: empty list, single node, or k=1 (no reversal needed)
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	// create a dummy node to handle edge cases
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		// find the kth node
		kth := getKthNode(prevGroupEnd.Next, k)
		if kth == nil {
			break // not enough nodes to form a gorup of size k
		}

		// Mark the start of current group and the node after kth
		groupStart := prevGroupEnd.Next
		nextGroupStart := kth.Next

		// Reverse the current group
		reverseList(groupStart, nextGroupStart)

		// Connect the reversed group back to the main list
		prevGroupEnd.Next = kth
		groupStart.Next = nextGroupStart

		// update prevGroupEnd for next iteration
		prevGroupEnd = groupStart
	}
	return dummy.Next
}

// getKthNode returns the kth node from the start node or nil if there are fewer than k nodes
func getKthNode(start *ListNode, k int) *ListNode {
	curr := start
	for i := 1; i < k && curr != nil; i++ {
		curr = curr.Next
	}
	return curr
}

// reverseList reverses a linked list from start to (but not including) end
func reverseList(start, end *ListNode) {
	prev := end
	curr := start

	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}
