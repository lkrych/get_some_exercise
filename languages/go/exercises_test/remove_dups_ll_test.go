import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesLL(t *testing.T) {
	input1 := []int{1, 2, 3, 3, 4, 4, 5}
	input2 := []int{5, 8, 6, 7, 8, 9, 7, 10}
	input3 := []int{100, 102, 101, 102, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	list1 = deleteDupsLL(list1)
	list2 = deleteDupsLL(list2)
	list3 = deleteDupsLL(list3)

	len1 := countLengthOfLL(list1)
	len2 := countLengthOfLL(list2)
	len3 := countLengthOfLL(list3)

	if len1 != 5 {
		t.Errorf("Expected deleteDups to make a list of length 5 not %v", len1)
		fmt.Println("Printing list1")
		printLL(list1)
	}

	if len2 != 6 {
		t.Errorf("Expected deleteDups to make a list of length 6 not %v", len2)
		fmt.Println("Printing list2")
		printLL(list2)
	}

	if len3 != 3 {
		t.Errorf("Expected deleteDups to make a list of length 3 not %v", len3)
		fmt.Println("Printing list3")
		printLL(list3)
	}
}

//if you created your own function, comment this out!
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

func countLengthOfLL(head *Node) int {
	count := 0
	currentNode := head
	for currentNode.next != nil {
		count++
		currentNode = currentNode.next
	}
	count++ //b/c you won't count the last one
	return count
}

func findValueInLL(head *Node, findVal int) bool {
	found := false
	currentNode := head
	for currentNode.next != nil {
		if currentNode.val == findVal {
			found = true
		}
		currentNode = currentNode.next
	}
	if currentNode.val == findVal { //catch the last node
		found = true
	}
	return found
}

func printLL(head *Node) {
	currentNode := head

	for currentNode.next != nil {
		fmt.Printf("  %v  ->", currentNode)
		currentNode = currentNode.next
	}
	fmt.Printf("  %v  \n", currentNode)
}
