package goExercises

import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
)

import (
	"fmt"
	"testing"
)

func TestCreatingAStackOfStacks(t *testing.T) {
	s := &stackOfStacks{
		capacity: 2,
	}
	for i := 1; i < 8; i++ {
		s.push(i)
	}

	if len(s.stacks) != 4 {
		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	testPop := s.pop()

	if testPop != 7 {
		fmt.Println(testPop)
		t.Errorf("Expected stackOfStacks pop to work like stack pop")
	}

	if len(s.stacks) != 3 {
		t.Errorf("Expected stack of stacks to have: 3 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	s.push(testPop)
	if len(s.stacks) != 4 {
		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	s.push(0)
	s.push(99)
	if len(s.stacks) != 5 {
		t.Errorf("Expected stack of stacks to have: 5 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

}
 
