import (
	"fmt"
	"testing"
)

func TestFindKthElement(t *testing.T) {
	input1 := []int{1, 2, 3, 3, 4, 4, 5}
	input2 := []int{5, 8, 6, 7, 8, 9, 7, 10}
	input3 := []int{100, 102, 101, 102, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	//remove _ if you are using the iterative solution
	list1Val, _ := findKthElement(list1, 2)
	list2Val, _ := findKthElement(list2, 5)
	list3Val, _ := findKthElement(list3, 3)

	if list1Val.val != 4 {
		t.Errorf("Expected kthElement to find the kth element 4 not %v", list1Val.val)
		fmt.Println("Printing list1")
		printLL(list1)
	}

	if list2Val.val != 7 {
		t.Errorf("Expected kthElement to find the kth element 7 not %v", list2Val.val)
		fmt.Println("Printing list2")
		printLL(list2)
	}

	if list3Val.val != 102 {
		t.Errorf("Expected kthElement to find the kth element 102 not %v", list3Val.val)
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

func printLL(head *Node) {
	currentNode := head

	for currentNode.next != nil {
		fmt.Printf("  %v  ->", currentNode)
		currentNode = currentNode.next
	}
	fmt.Printf("  %v  \n", currentNode)
}
