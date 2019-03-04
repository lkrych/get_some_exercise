import "testing"

func TestCreatingAStack(t *testing.T) {
	s := &stack{
		data: []int{1, 2, 3, 4, 5, 6},
	}
	testPop := s.pop()
	testPop2 := s.pop()

	if len(s.data) != 4 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop)
	}
	if s.peek() != 4 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 4 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements", len(s.data))
	}

	s.push(testPop)
	if len(s.data) != 5 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop2)
	}
	if s.peek() != 6 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 5 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
	}

	s.push(testPop2)
	if len(s.data) != 6 {
		t.Errorf("Expected push to return elements to the top of the stack")
	}
	if s.peek() != 5 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 6 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
	}

	for i := 0; i < 6; i++ {
		s.pop()
	}
	if !s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack should be empty")
	}
}
