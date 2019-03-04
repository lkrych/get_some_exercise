func deleteDupsLL(head *Node) *Node {
	intMap := map[int]bool{}
	prev := &Node{}
	currentNode := head
	for currentNode.next != nil {
		if _, ok := intMap[currentNode.val]; ok {
			prev.next = currentNode.next
		} else {
			intMap[currentNode.val] = true
			prev = currentNode
		}
		currentNode = currentNode.next
	}
	if _, ok := intMap[currentNode.val]; ok {
		prev.next = currentNode.next
	}
	return head
}