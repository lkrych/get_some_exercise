import "testing"

func TestCreatingAMinStack(t *testing.T) {
	s := &stackWithMin{
		data: []int{},
	}
	for i := 1; i < 7; i++ {
		s.push(i)
	}
	testPop := s.pop()
	testPop2 := s.pop()

	if len(s.data) != 4 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop)
	}
	if s.min() != 1 {
		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
	}

	s.push(testPop)
	if len(s.data) != 5 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop2)
	}
	if s.min() != 1 {
		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
	}

	s.push(testPop2)
	if len(s.data) != 6 {
		t.Errorf("Expected push to return elements to the top of the stack")
	}

	if s.min() != 1 {
		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
	}

	s.push(0)
	if len(s.data) != 7 {
		t.Errorf("Expected push to return elements to the top of the stack")
	}

	if s.min() != 0 {
		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
	}

}
