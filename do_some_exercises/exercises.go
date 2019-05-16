package goExercises

//return true if the letter a is within two spaces of the letter z
func nearbyAZ(word string) bool {
	indexA := -1000
	for i := 0; i < len(word); i++ {
		if string(word[i]) == "a" {
			indexA = i
		}

		if string(word[i]) == "z" {
			if i-indexA <= 3 {
				return true
			}
		}
	}
	return false
}
