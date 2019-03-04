import (
	"fmt"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{5, 6, 7, 8, 9, 10}
	input3 := []int{100, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	middleNode1 := findNodeInLL(list1, 3)
	middleNode2 := findNodeInLL(list2, 7)
	middleNode3 := findNodeInLL(list3, 101)

	if middleNode1.val != 3 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	if middleNode2.val != 7 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	if middleNode3.val != 101 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	deleteMiddleNode(middleNode1)
	deleteMiddleNode(middleNode2)
	deleteMiddleNode(middleNode3)

	len1 := countLengthOfLL(list1)
	len2 := countLengthOfLL(list2)
	len3 := countLengthOfLL(list3)

	if len1 != 4 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 4 not %v", len1)
		fmt.Println("Printing d4")
		printLL(list1)
	}
	if findValueInLL(list1, 3) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 3")
		fmt.Println("Printing list1")
		printLL(list1)
	}

	if len2 != 5 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 5 not %v", len2)
		fmt.Println("Printing list2")
		printLL(list2)
	}
	if findValueInLL(list2, 7) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 7")
		fmt.Println("Printing list2")
		printLL(list2)
	}

	if len3 != 2 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 2 not %v", len3)
		fmt.Println("Printing list3")
		printLL(list3)
	}
	if findValueInLL(list3, 101) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 101")
		fmt.Println("Printing list3")
		printLL(list3)
	}

}

func findNodeInLL(head *Node, val int) *Node {
	for head.next != nil {
		if head.val == val {
			return head
		}
		head = head.next
	}
	return nil
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