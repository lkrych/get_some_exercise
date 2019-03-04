func TestPassPhrase(t *testing.T) {
	test1 := passPhrase("doody goody woody")
	test2 := passPhrase("doody doody woody")
	test3 := passPhrase("woody goody woody")
	test4 := passPhrase("It is not I who is the Creature!")
	test5 := passPhrase("Hello, my name is Bill.")
	test6 := passPhraseAdvent()

	if test1 != true {
		t.Errorf("Expected answer to be true, not %v", test1)
	}

	if test2 != false {
		t.Errorf("Expected answer to be false, not %v", test2)
	}

	if test3 != false {
		t.Errorf("Expected answer to be false, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected answer to be false, not %v", test4)
	}

	if test5 != true {
		t.Errorf("Expected answer to be true, not %v", test5)
	}

	if test6 != 325 {
		t.Errorf("Expected answer to be 325, not %v", test6)
	}
}

// func TestPassPhraseAnagram(t *testing.T) {
// 	test1 := passPhraseAnagram("abcde fghij")
// 	test2 := passPhraseAnagram("abcde xyz ecdab")
// 	test3 := passPhraseAnagram("a ab abc abd abf abj")
// 	test4 := passPhraseAnagram("iiii oiii ooii oooi oooo")
// 	test5 := passPhraseAnagram("oiii ioii iioi iiio")
// 	test6 := passPhraseAnagramAdvent()

// 	if test1 != true {
// 		t.Errorf("Expected answer to be true, not %v", test1)
// 	}

// 	if test2 != false {
// 		t.Errorf("Expected answer to be false, not %v", test2)
// 	}

// 	if test3 != true {
// 		t.Errorf("Expected answer to be true, not %v", test3)
// 	}

// 	if test4 != true {
// 		t.Errorf("Expected answer to be false, not %v", test4)
// 	}

// 	if test5 != false {
// 		t.Errorf("Expected answer to be false, not %v", test5)
// 	}

// 	if test6 != 6 {
// 		t.Errorf("Expected answer to be 6, not %v", test6)
// 	}
// }
