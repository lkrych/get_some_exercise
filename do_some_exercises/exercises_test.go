package goExercises

import (
	"fmt"
	"testing"
)

func TestCreateALinkedList(t *testing.T) {
	input1 := []int{10, 5, 8, 1, 3, 2, 4}
	input2 := []int{4, 7, 2, 3, 1}
	input3 := []int{105, 783, 98, 77, 33, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	partition1 := partitionLL(list1, 5)
	partition2 := partitionLL(list2, 3)
	partition3 := partitionLL(list3, 105)

	if findIdxOfVal(partition1, 4) > 5 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition1)
	}

	if findIdxOfVal(partition1, 10) < 4 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition1)
	}

	if findIdxOfVal(partition2, 7) < 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition2)
	}

	if findIdxOfVal(partition1, 2) > 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition2)
	}

	if findIdxOfVal(partition3, 102) > 5 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition3)
	}

	if findIdxOfVal(partition3, 783) < 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition3)
	}

}

func findIdxOfVal(head *Node, search int) int {
	idx := 0
	for head.val != search {
		idx++
		head = head.next
	}
	return idx
}

func createALinkedList(input []int) *Node {
	head := &Node{
		val:  input[0],
		next: nil,
	}
	currentNode := head
	for i := 1; i < len(input); i++ {
		nextNode := &Node{
			val:  input[i],
			next: nil,
		}
		currentNode.next = nextNode
		currentNode = nextNode
	}
	return head
}

func printLL(head *Node) {
	currentNode := head

	for currentNode.next != nil {
		fmt.Printf("  %v  ->", currentNode)
		currentNode = currentNode.next
	}
	fmt.Printf("  %v  \n", currentNode)
}
