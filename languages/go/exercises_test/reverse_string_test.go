func TestReverse(t *testing.T) {
	word := reverse("candelabra")
	sentence := reverse("reverse this sentence")

	if word != "arbalednac" {
		t.Errorf("Expected reverse to be able to reverse a string")
	}

	if sentence != "ecnetnes siht esrever" {
		t.Errorf("Expected reverse to be able to reverse a sentence, but got %v", sentence)
	}

}