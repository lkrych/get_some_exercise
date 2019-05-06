package goExercises

type Node struct {
	val  int
	next *Node
}

//write code to partition a linked list around a value x, such that all nodes less than
//x come before all nodes greater than or equal to x
func partitionLL(head *Node, partition int) *Node {
	beforeStart := &Node{}
	beforeEnd := &Node{}
	afterStart := &Node{}
	afterEnd := &Node{}

	//partition
	currentNode := head
	for currentNode != nil {
		next := currentNode.next
		currentNode.next = nil
		if currentNode.val < partition {
			//insert node into end of before list
			if beforeStart.val == 0 {
				beforeStart = currentNode
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = currentNode
				beforeEnd = currentNode
			}
		} else {
			//insert node into end of after list
			if afterStart.val == 0 {
				afterStart = currentNode
				afterEnd = afterStart
			} else {
				afterEnd.next = currentNode
				afterEnd = currentNode
			}
		}
		currentNode = next
	}
	if beforeStart.val == 0 {
		return afterStart
	}

	//merge the two lists
	beforeEnd.next = afterStart
	return beforeStart
}
