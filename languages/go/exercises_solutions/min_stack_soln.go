// type stackWithMin struct {
// 	data     []int
// 	minStack stack
// }

// func (s *stackWithMin) pop() int {
// 	popped := s.data[len(s.data)-1]
// 	s.data = s.data[:len(s.data)-1]
// 	if popped == s.min() { //if you've popped the min, move it off the minStack!
// 		s.minStack.pop()
// 	}
// 	return popped
// }

// func (s *stackWithMin) push(item int) {
// 	s.data = append(s.data, item)
// 	if item <= s.min() {
// 		s.minStack.push(item)
// 	}
// }

// func (s *stackWithMin) min() int {
// 	if s.minStack.isEmpty() {
// 		return 100000000000000000 //really large number
// 	}
// 	return s.minStack.peek()
// }
