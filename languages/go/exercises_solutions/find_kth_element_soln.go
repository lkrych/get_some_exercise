//there are two ways to solve this problem: a recursive way and an iterative way
// the iterative way is more efficient, the recursive way is easier to read
//
// recursive:
//
// func findKthElement(head *Node, k int) (*Node, int) {
// 	if head.next == nil {
// 		return head, 1
// 	}
// 	var node, ith = findKthElement(head.next, k)
// 	ith++
// 	if ith == k {
// 		return head, ith
// 	}
// 	return node, ith
// }
//
// iterative:
//
// func findKthElement(head *Node, k int) *Node {
// 	p1 := head
// 	p2 := head

// 	//move p1 k nodes into the list
// 	for i := 0; i < k; i++ {
// 		if p1 == nil {
// 			return nil
// 		}
// 		p1 = p1.next
// 	}

// 	//move them at the same pace. when p1 hits the end, p2 will be at the right ele
// 	for p1 != nil {
// 		p1 = p1.next
// 		p2 = p2.next
// 	}
// 	return p2
// }