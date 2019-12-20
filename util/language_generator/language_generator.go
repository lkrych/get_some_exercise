package main

import (
	"bufio"
	"fmt"
	"get_some_exercise/util"
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
	re := regexp.MustCompile(`\r?\n`)
	name = re.ReplaceAllString(name, "")
	util.CheckErr(err)
	for {
		util.PrintSpace()
		fmt.Printf("Are you sure you want to call your exercise: %v y/n ", name)
		yesOrNo, err := r.ReadString('\n')
		util.CheckErr(err)

		match, err := regexp.MatchString("y|n", yesOrNo)
		if match {
			match, err = regexp.MatchString("y", yesOrNo)
			if match {
				break
			}
		} else {
			util.FormatPrint("You need to indicate y or n. Please try again.")
			fmt.Print("What is the name of your exercise? (exclude the file suffix) ")
			name, err = r.ReadString('\n')
		}
	}
	return name
}

func createNewFiles(name string, r *bufio.Reader) {

	for {
		util.FormatPrint("Does this exercise already exist in the go language folder? y/n ")
		exists, err := r.ReadString('\n')

		if err != nil {
			util.FormatPrint("There was an error reading your input, try again.")
		}

		match, err := regexp.MatchString("y|n", exists)
		if match {
			match, err = regexp.MatchString("y", exists)
			if match {
				createEverythingButGo(name)
				break
			} else {
				createEverything(name)
				break
			}
		} else {
			util.FormatPrint("You need to indicate y or n. Please try again.")
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
		suffix := util.GetFileSuffix(lang)
		execsFile := fmt.Sprintf("../../languages/%v/exercises_itemized/%v%v", lang, name, suffix)
		execs, err := os.Create(execsFile)
		util.CheckErr(err)
		execs.Close()
		testFile := fmt.Sprintf("../../languages/%v/exercises_test/%v_test%v", lang, name, suffix)
		tests, err := os.Create(testFile)
		util.CheckErr(err)
		tests.Close()
		solnsFile := fmt.Sprintf("../../languages/%v/exercises_solutions/%v_soln%v", lang, name, suffix)
		solns, err := os.Create(solnsFile)
		util.CheckErr(err)
		solns.Close()
	}
}
