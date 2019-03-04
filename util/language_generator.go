package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	createNewFiles()
}

func createNewFiles() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Does this exercise already exist in the go language folder? y/n")
		exists, err := reader.ReadString('\n')

		if err != nil {
			formatPrint("There was an error reading your input, try again.")
		}

		match := regexp.MatchString("y|n", exists)
		if match {
			break
		} else {
			formatPrint("You need to indicate y or n. Please try again.")
		}
	}

}
