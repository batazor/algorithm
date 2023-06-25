package loop

type Node struct {
	Value int
	Next  *Node
}

func IsLoop_Floyd(head *Node) (int, bool) {
	if head == nil {
		return 0, false
	}

	slow, fast := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return 0, false
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Value, true
}
