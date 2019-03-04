func TestSumLists(t *testing.T) {
	i1 := []int{1, 2, 3}
	i2 := []int{5, 8, 6}
	sum1 := []int{6, 0, 0, 1}

	i3 := []int{3, 5}
	i4 := []int{1, 8}
	sum2 := []int{4, 3, 1}

	i5 := []int{6, 4, 5}
	i6 := []int{1, 1, 1}
	sum3 := []int{7, 5, 6}

	list1 := createALinkedList(i1)
	list2 := createALinkedList(i2)
	answer1 := listToArray(createALinkedList(sum1))

	list3 := createALinkedList(i3)
	list4 := createALinkedList(i4)
	answer2 := listToArray(createALinkedList(sum2))

	list5 := createALinkedList(i5)
	list6 := createALinkedList(i6)
	answer3 := listToArray(createALinkedList(sum3))

	returned1 := listToArray(sumLists(list1, list2))
	returned2 := listToArray(sumLists(list3, list4))
	returned3 := listToArray(sumLists(list5, list6))

	if !cmp.Equal(answer1, returned1) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list1, list2, answer1, returned1)

	}

	if !cmp.Equal(answer2, returned2) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list3, list4, answer2, returned2)
	}

	if !cmp.Equal(answer3, returned3) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list5, list6, answer3, returned3)
	}
}

func listToArray(head *Node) []int {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.val)
		head = head.next
	}
	return vals
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