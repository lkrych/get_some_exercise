import "testing"

func TestCreatingAQueueOfStacks(t *testing.T) {
	q := &myQueue{}
	for i := 1; i < 7; i++ {
		q.add(i)
	}

	testremove := q.remove()
	testremove2 := q.remove()

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	q.add(testremove)

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	q.add(testremove2)

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	for i := 0; i < 6; i++ {
		q.remove()
	}
	if !q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue should be empty")
	}
}
