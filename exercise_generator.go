package main

import (
	"bufio"
	"fmt"
	"get_some_exercise/util"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//clean up old exercises if they exist
	err := util.RemoveOldExercises()
	util.CheckErr(err)
	fmt.Print("🏃‍♀️ Get Some Exercise 🏃\n")
	util.PrintSpace()
	// ask for mode
	mode := askForMode()
	if mode == "random" {
		//ask for number of Exercises
		numExercises := askForExercises()
		//ask for language
		language := askForLanguage()
		//create quiz folder with exercises and tests
		makeExercises(numExercises, language)
	} else if mode == "targeted" {
		//ask for language
		language := askForLanguage()
		//ask for exercise
		exercise := askForExercise(language)
		//make excercise
		makeExercise(exercise, language)
	}
	fmt.Print("🏃‍♀️ Your files have been written in ./do_some_exercises. It's time to get running 🏃\n")
}

type filePaths struct {
	exerciseFile string
	testFile     string
	solnsFile    string
}

//use stdinput to ask for random or targeted questions
func askForMode() string {
	mode := "random" // default to random

	for { //iterate until mode is picked
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Targeted: Practice a specific exercise \n")
		fmt.Printf("Random: Pick number of exercises and language\n")
		util.PrintSpace()
		fmt.Print("Enter mode you want (targeted or random): ")
		mode, _ = reader.ReadString('\n')

		//strip whitespace from user input
		mode = strings.TrimSpace(mode)

		match, _ := regexp.MatchString("random|targeted", mode)
		if match {
			break
		} else {
			util.FormatPrint("You need to pick either random or targeted")
			fmt.Printf("You printed %v", mode)
		}
	}
	return mode
}

//use stdinput to ask for index to exercise
func askForExercise(language string) int {
	numExercise := 0 //default to first
	exercisesPath := fmt.Sprintf("./languages/%v/exercises_itemized", language)
	files, err := ioutil.ReadDir(exercisesPath)
	util.CheckErr(err)

	//print out all the files
	for idx, file := range files {
		fmt.Printf("[%d]: %v \n", idx, file.Name())
	}

	for { //iterate until the number is between bounds
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter number of exercise that you want to practice: ")
		numExerciseString, err := reader.ReadString('\n')
		numExercise, err = strconv.Atoi(strings.Split(numExerciseString, "\n")[0])

		if err != nil {
			util.FormatPrint("That isn't a valid number!")
		}

		if numExercise < 0 || numExercise > len(files)-1 {
			util.FormatPrint("You need to pick a valid exercise!")
		} else {
			break
		}
	}
	return numExercise
}

//use stdinput to ask for a number of Exercises to practice
func askForExercises() int {
	numExercises := 0
	for { //iterate until the number is between bounds
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter number of exercises that you want to practice (min: 1, max: 10): ")
		numExercisesString, err := reader.ReadString('\n')
		numExercises, err = strconv.Atoi(strings.Split(numExercisesString, "\n")[0])

		if err != nil {
			util.FormatPrint("That isn't a valid number!")
		}

		if numExercises < 1 || numExercises > 10 {
			util.FormatPrint("You need to pick an integer between 1 and 10")
		} else {
			break
		}
	}
	return numExercises
}

//use stdinnput to ask for language of Exercises to practice
func askForLanguage() string {
	language := "go" // default to go

	for { //iterate until language is picked
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter language you want to practice (python, go, c, cpp, elixir, ocaml, javascript): ")
		language, _ = reader.ReadString('\n')

		//strip whitespace from user input
		language = strings.TrimSpace(language)

		match, _ := regexp.MatchString("python|go|c|cpp|elixir|ocaml|javascript|", language)
		if match {
			break
		} else {
			util.FormatPrint("You need to pick either python, go, c, cpp, javascript, elixir, or ocaml")
			fmt.Printf("You printed %v", language)
		}
	}
	return language
}

func makeExercise(numExercise int, language string) {
	//get fileSuffix
	fileSuffix := util.GetFileSuffix(language)

	//create directory and files
	filePaths := createFiles(fileSuffix)

	//read exercises
	exercisesPath := fmt.Sprintf("./languages/%v/exercises_itemized", language)
	files, err := ioutil.ReadDir(exercisesPath)
	util.CheckErr(err)

	//Open exercise and test files for writing

	exerciseFile, err := os.OpenFile(filePaths.exerciseFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer exerciseFile.Close()
	testFile, err := os.OpenFile(filePaths.testFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer testFile.Close()
	solnFile, err := os.OpenFile(filePaths.solnsFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer solnFile.Close()

	//make header
	headerFile, err := os.Create("./do_some_exercises/exercises.h")
	util.CheckErr(err)
	defer headerFile.Close()

	initFiles(language, exerciseFile, testFile, solnFile)

	// add exercises and tests to file
	fileName := files[numExercise].Name()

	//read exercise file
	exercise, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_itemized/%v", language, fileName))
	util.CheckErr(err)

	// read exercise test file
	split := strings.Split(fileName, ".")
	testName := fmt.Sprintf("%v_test%v", split[0], fileSuffix)
	test, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_test/%v", language, testName))
	util.CheckErr(err)

	//read exercise soln file
	solnName := fmt.Sprintf("%v_soln%v", split[0], fileSuffix)
	soln, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_solutions/%v", language, solnName))

	//if language is cpp write exercise
	if language == "cpp" {
		split := strings.Split(fileName, ".")
		signatureName := fmt.Sprintf("%v%v", split[0], ".h")
		functionSignature, err := ioutil.ReadFile(fmt.Sprintf("./languages/cpp/exercises_headers/%v", signatureName))
		util.CheckErr(err)
		_, err = headerFile.Write([]byte(functionSignature))
		util.CheckErr(err)
	}

	//check to make sure files are being read correctly
	fmt.Println("Exercises and Tests for", fileName)

	_, err = exerciseFile.Write(exercise)
	util.CheckErr(err)
	_, err = exerciseFile.Write([]byte("\n \n"))
	util.CheckErr(err)
	_, err = testFile.Write(test)
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\n \n"))
	util.CheckErr(err)
	_, err = solnFile.Write(soln)
	util.CheckErr(err)
	_, err = solnFile.Write([]byte("\n \n"))
	util.CheckErr(err)

	switch language {
	case "c":
		finishTestFileC(testFile)
	case "cpp":
		finishTestFileCpp(testFile)
	case "elixir":
		finishTestFileElixir(testFile)
	case "python":
		finishTestFilePython(testFile)
	}
	createHelperFile(fileSuffix)
	createMakeFile()
}

func makeExercises(numExercises int, language string) {

	//get fileSuffix
	fileSuffix := util.GetFileSuffix(language)

	//create directory and files
	filePaths := createFiles(fileSuffix)

	//read exercises
	exercisesPath := fmt.Sprintf("./languages/%v/exercises_itemized", language)
	files, err := ioutil.ReadDir(exercisesPath)
	util.CheckErr(err)

	//grab random indices corresponding to files in files slice
	randIdxs := randomIndices(numExercises, len(files))

	//Open exercise and test files for writing

	exerciseFile, err := os.OpenFile(filePaths.exerciseFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer exerciseFile.Close()
	testFile, err := os.OpenFile(filePaths.testFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer testFile.Close()
	solnFile, err := os.OpenFile(filePaths.solnsFile, os.O_RDWR, 0666)
	util.CheckErr(err)
	defer solnFile.Close()

	//make header
	headerFile, err := os.Create("./do_some_exercises/exercises.h")
	util.CheckErr(err)
	defer headerFile.Close()

	initFiles(language, exerciseFile, testFile, solnFile)

	// add exercises and tests to file
	for _, idx := range randIdxs {
		fileName := files[idx].Name()

		//read exercise file
		exercise, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_itemized/%v", language, fileName))
		util.CheckErr(err)

		// read exercise test file
		split := strings.Split(fileName, ".")
		testName := fmt.Sprintf("%v_test%v", split[0], fileSuffix)
		test, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_test/%v", language, testName))
		util.CheckErr(err)

		//read exercise soln file
		solnName := fmt.Sprintf("%v_soln%v", split[0], fileSuffix)
		soln, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_solutions/%v", language, solnName))

		//if language is cpp write exercise
		if language == "cpp" {
			split := strings.Split(fileName, ".")
			signatureName := fmt.Sprintf("%v%v", split[0], ".h")
			functionSignature, err := ioutil.ReadFile(fmt.Sprintf("./languages/cpp/exercises_headers/%v", signatureName))
			util.CheckErr(err)
			_, err = headerFile.Write([]byte(functionSignature))
			util.CheckErr(err)
		}

		//check to make sure files are being read correctly
		fmt.Println("Exercises and Tests for", fileName)

		_, err = exerciseFile.Write(exercise)
		util.CheckErr(err)
		_, err = exerciseFile.Write([]byte("\n \n"))
		util.CheckErr(err)
		_, err = testFile.Write(test)
		util.CheckErr(err)
		_, err = testFile.Write([]byte("\n \n"))
		util.CheckErr(err)
		_, err = solnFile.Write(soln)
		util.CheckErr(err)
		_, err = solnFile.Write([]byte("\n \n"))
		util.CheckErr(err)
	}

	switch language {
	case "c":
		finishTestFileC(testFile)
	case "cpp":
		finishTestFileCpp(testFile)
	case "elixir":
		finishTestFileElixir(testFile)
	case "python":
		finishTestFilePython(testFile)
	}
	createHelperFile(fileSuffix)
	createMakeFile()
}

func createFiles(suffix string) *filePaths {
	os.Mkdir("./do_some_exercises", os.ModePerm)
	execsFile := fmt.Sprintf("./do_some_exercises/exercises%v", suffix)
	execs, err := os.Create(execsFile)
	util.CheckErr(err)
	execs.Close()
	testFile := fmt.Sprintf("./do_some_exercises/exercises_test%v", suffix)
	tests, err := os.Create(testFile)
	util.CheckErr(err)
	tests.Close()
	solnsFile := fmt.Sprintf("./do_some_exercises/exercises_solns%v", suffix)
	solns, err := os.Create(solnsFile)
	util.CheckErr(err)
	solns.Close()

	return &filePaths{
		exerciseFile: execsFile,
		testFile:     testFile,
		solnsFile:    solnsFile,
	}
}

func randomIndices(n int, arrLen int) []int {
	//a map is a more efficient way to enforce unique idxs
	m := make(map[int]bool)
	//seed
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	idxCount := 0
	for idxCount < n {
		//generate a random int
		newIdx := r1.Intn(arrLen)
		//check if that int has been found
		if _, ok := m[newIdx]; ok {
			//if it has, continue looping
			continue
		}
		//otherwise add it to the map and increase the idx count
		m[newIdx] = true
		idxCount++
	}
	//create arr of idxs
	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	fmt.Printf("The indexes we are using are %v \n", keys)
	return keys
}

//initialize files

func initFiles(language string, exerciseFile, testFile, solnFile *os.File) {
	switch language {
	case "python":
		initFilesPython(exerciseFile, testFile, solnFile)
	case "go":
		initFilesGo(exerciseFile, testFile, solnFile)
	case "c":
		initFilesC(exerciseFile, testFile, solnFile)
	case "cpp":
		initFilesCpp(exerciseFile, testFile, solnFile)
	case "elixir":
		initFilesElixir(exerciseFile, testFile, solnFile)
	case "ocaml":
		initFilesOcaml(exerciseFile, testFile, solnFile)
	case "javascript":
		initFilesJavascript(exerciseFile, testFile, solnFile)
	}
}

func createHelperFile(fileSuffix string) {
	//create helper for python or header for cpp
	if fileSuffix == ".py" {
		src := fmt.Sprintf("./util/helper/helper%v", fileSuffix)
		dst := fmt.Sprintf("./do_some_exercises/helper%v", fileSuffix)
		cpCmd := exec.Command("cp", "-rf", src, dst)
		err := cpCmd.Run()
		if err != nil {
			util.CheckErr(err)
		}
	}
}

//steps to initialize files differ by language
func initFilesGo(exerciseFile, testFile, solnFile *os.File) {
	//init files with package main
	_, err := exerciseFile.Write([]byte("package goExercises\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("package goExercises\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("import (\n\"fmt\"\n\"testing\"\n\"github.com/google/go-cmp/cmp\"\n)\n\n"))
	util.CheckErr(err)
	_, err = solnFile.Write([]byte("package goExercises\n\n"))
	util.CheckErr(err)
}

func initFilesPython(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("import unittest\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("from exercises import *\n\n"))
	util.CheckErr(err)
}

func initFilesOcaml(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("open OUnit2;;\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("open Exercises;;\n\n"))
	util.CheckErr(err)
}

func initFilesC(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("#include \"exercises.c\"\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("#include <CUnit/CUnit.h>\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("#include <CUnit/Basic.h>\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("int init_suite(void) { return 0; }\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("int clean_suite(void) { return 0; }\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("/************* Your Code goes here **************/\n\n"))
	util.CheckErr(err)
}

func initFilesCpp(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("#include \"gtest/gtest.h\"\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("#include \"exercises.h\"\n"))
	util.CheckErr(err)
}

func initFilesElixir(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("#Start ExUnit the testrunner\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("ExUnit.start()\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("#create a new test module (test case) and use ExUnit.Case\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("defmodule TestExercises do\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tuse ExUnit.Case, async: true\n\n"))
	util.CheckErr(err)
}

func initFilesJavascript(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("var test = require('tape'); \n"))
	util.CheckErr(err)
}

func finishTestFileC(testFile *os.File) {
	_, err := testFile.Write([]byte("/************* Test Runner Code goes here **************/\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("int main ( void )\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("{\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tCU_pSuite pSuite = NULL;\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t/* initialize the CUnit test registry */ \n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tif ( CUE_SUCCESS != CU_initialize_registry() ) \n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t\treturn CU_get_error();\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t/* add a suite to the registry */\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tpSuite = CU_add_suite( \"exercises_test_suite\", init_suite, clean_suite ); \n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tif ( NULL == pSuite ) { \n\t\tCU_cleanup_registry();\n\t\treturn CU_get_error();\n\t}\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t/* add the tests to the suite */\n\tif ( (NULL == CU_add_test(pSuite, \"test_fib\", test_fib)) /* || \n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t\t(NULL == CU_add_test(pSuite, \"max_test_2\", max_test_2)) || \n\t\t(NULL == CU_add_test(pSuite, \"max_test_3\", max_test_3)) */ \n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t)\n{\n\tCU_cleanup_registry(); \n\treturn CU_get_error();\n}\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\t// Run all tests using the basic interface\n/tCU_basic_set_mode(CU_BRM_VERBOSE);\n/tCU_basic_run_tests();\n/tprintf(\"\n\");\n/tCU_basic_show_failures(CU_get_failure_list());\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tprintf(\"\n\n\"); \n\n\t/* Clean up registry and return */\n\tCU_cleanup_registry();\n/treturn CU_get_error();\n\n}"))
	util.CheckErr(err)

}

func finishTestFileCpp(testFile *os.File) {
	_, err := testFile.Write([]byte("\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("int main(int argc, char **argv) {\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\ttesting::InitGoogleTest(&argc, argv);\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\treturn RUN_ALL_TESTS();\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("}\n"))
	util.CheckErr(err)
}

func finishTestFilePython(testFile *os.File) {
	_, err := testFile.Write([]byte("\n\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("if __name__ == '__main__':\n"))
	util.CheckErr(err)
	_, err = testFile.Write([]byte("\tunittest.main()"))
	util.CheckErr(err)
}

func finishTestFileElixir(testFile *os.File) {
	_, err := testFile.Write([]byte("end"))
	util.CheckErr(err)
}

func createMakeFile() {
	makeFilePath := fmt.Sprintf("./do_some_exercises/Makefile")
	makeFile, err := os.Create(makeFilePath)
	util.CheckErr(err)
	defer makeFile.Close()

	_, err = makeFile.Write([]byte("go:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tgo test \n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("javascript:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tnode -r esm exercises_test.js | ../node_modules/.bin/tap-dot \n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("python:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tpython -m unittest exercises_test.py  \n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("c:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tgcc exercises_test.c -lcunit -o test ./test \n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("cpp:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tg++ exercises_test.cpp exercises.cpp -lgtest -lgtest_main -pthread -std=c++11 -o test \n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\t./test\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("elixir:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\telixir -r exercises.ex exercises_test.exs\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("ocaml_build:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\tocamlbuild -pkgs OUnit exercises_test.byte\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("ocaml_run:\n"))
	util.CheckErr(err)
	_, err = makeFile.Write([]byte("\t./exercises_test.byte\n"))
	util.CheckErr(err)
}
