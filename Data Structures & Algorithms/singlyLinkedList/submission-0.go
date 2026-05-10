type LinkedList struct {
	head *Node

}

type Node struct {
	data int
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (ll *LinkedList) Get(index int) int {

	current := ll.head
	for i := 0; i < index && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		return -1
	}

	return current.data
}

func (ll *LinkedList) InsertHead(val int) {
	temp := ll.head
	newHead := &Node{data: val, next: temp}
	ll.head = newHead
}

func (ll *LinkedList) InsertTail(val int) {

	newTail := &Node{data: val, next: nil}

	// Edge case: empty list
	if ll.head == nil {
		ll.head = newTail
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newTail
}

func (ll *LinkedList) Remove(index int) bool {

    // Edge case: empty list
    if ll.head == nil {
        return false
    }

    // Special case: remove the head
    if index == 0 {
        ll.head = ll.head.next
        return true
    }

	// Traverse until prev = node before the target
    prev := ll.head

	for i := 0; i < (index - 1) && prev.next != nil; i++ {
		prev = prev.next
	}

    // If prev.next is nil, target is out of bounds
    if prev.next == nil {
        return false
    }

    // Bypass the target: prev jumps over it
    prev.next = prev.next.next
    return true
}

func (ll *LinkedList) GetValues() []int {
	var values []int

	current := ll.head
	for i := 0 ; current != nil; i++ {
		values = append(values, current.data)
		current = current.next
	}

	return values
}
