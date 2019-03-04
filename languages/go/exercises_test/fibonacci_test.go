
func testFibonacci(t *testing.T) {
	t1, err1 := fibonacci(3)
	t2, err2 := fibonacci(6)
	t3, err3 := fibonacci(7)
	t4, err4 := fibonacci(9)
	t5, err5 := fibonacci(15)
	_, err6 := fibonacci(-1)
	_, err7 := fibonacci(0)
	ans1 := []int{1, 1, 2}
	ans2 := []int{1, 1, 2, 3, 5, 8}
	ans3 := []int{1, 1, 2, 3, 5, 8, 13}
	ans4 := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	ans5 := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	errorMessage := "n must be greater than 0"

	if !cmp.Equal(t1, ans1) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 3, ans1, t1)
	}

	if err1 != nil {
		t.Errorf("There was not supposed to be an error for t1")
	}

	if !cmp.Equal(t2, ans2) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 6, ans2, t2)
	}

	if err2 != nil {
		t.Errorf("There was not supposed to be an error for t2")
	}

	if !cmp.Equal(t3, ans3) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 7, ans3, t3)
	}

	if err3 != nil {
		t.Errorf("There was not supposed to be an error for t3")
	}

	if !cmp.Equal(t4, ans4) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 9, ans4, t4)
	}

	if err4 != nil {
		t.Errorf("There was not supposed to be an error for t4")
	}

	if !cmp.Equal(t5, ans5) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 15, ans5, t5)
	}
	if err5 != nil {
		t.Errorf("There was not supposed to be an error for t5")
	}

	if err6.Error() != errorMessage {
		t.Errorf("Expected the function to throw an error for input less than 1")
	}

	if err7.Error() != errorMessage {
		t.Errorf("Expected the function to throw an error for input less than 1")
	}
}