// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// )

// func passPhrase(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		if _, ok := wordMap[word]; ok {
// 			fmt.Printf("%v is the repeated word in %v \n", word, passphrase)
// 			return false
// 		}
// 		wordMap[word] = 1
// 	}
// 	return true
// }

// func passPhraseAdvent() int {
// 	phraseCount := 0
// 	content, err := ioutil.ReadFile("./advent_input/passphrase_input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	splitByLine := strings.Split(string(content), "\n")
// 	for _, line := range splitByLine {
// 		if passPhrase(line) {
// 			phraseCount++
// 		}
// 	}
// 	return phraseCount
// }

// func passPhraseAnagram(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		sorted := SortString(word)
// 		if _, ok := wordMap[sorted]; ok {
// 			return false
// 		}
// 		wordMap[sorted] = 1
// 	}
// 	return true
// }