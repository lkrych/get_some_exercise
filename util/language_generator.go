package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := getName(reader)
	createNewFiles(name, reader)
}

func getName(r *bufio.Reader) string {
	fmt.Print("What is the name of your exercise? (exclude the file suffix) ")
	name, err := r.ReadString('\n')
	CheckErr(err)
	for {
		printSpace()
		fmt.Printf("Are you sure you want to call your exericise: %v y/n ", name)
		yesOrNo, err := r.ReadString('\n')
		CheckErr(err)

		match, err := regexp.MatchString("y|n", yesOrNo)
		if match {
			match, err = regexp.MatchString("y", yesOrNo)
			if match {
				break
			}
		} else {
			FormatPrint("You need to indicate y or n. Please try again.")
			fmt.Print("What is the name of your exercise? (exclude the file suffix) ")
			name, err = r.ReadString('\n')
		}
	}
	return name
}

func createNewFiles(name string, r *bufio.Reader) {

	for {
		fmt.Print("Does this exercise already exist in the go language folder? y/n")
		exists, err := r.ReadString('\n')

		if err != nil {
			FormatPrint("There was an error reading your input, try again.")
		}

		match, err := regexp.MatchString("y|n", exists)
		if match {
			match, err = regexp.MatchString("y", exists)
			if match {
				createEverythingButGo(name)
			} else {
				createEverything(name)
			}
		} else {
			FormatPrint("You need to indicate y or n. Please try again.")
		}
	}
}

func createEverything(name string) {
	createFiles(name, []string{"go", "python", "c", "elixir"})
}

func createEverythingButGo(name string) {
	createFiles(name, []string{"python", "c", "elixir"})
}

func createFiles(name string, langArray []string) {
	for _, lang := range langArray {
		suffix := GetFileSuffix(lang)
		execsFile := fmt.Sprintf("../languages/%v/exercises_itemized/%v.%v", lang, name, suffix)
		execs, err := os.Create(execsFile)
		CheckErr(err)
		execs.Close()
		testFile := fmt.Sprintf("../languages/%v/exercises_test/%v_test.%v", lang, name, suffix)
		tests, err := os.Create(testFile)
		CheckErr(err)
		tests.Close()
		solnsFile := fmt.Sprintf("../languages/%v/exercises_solutions/%v_soln.%v", lang, name, suffix)
		solns, err := os.Create(solnsFile)
		CheckErr(err)
		solns.Close()
	}
}
