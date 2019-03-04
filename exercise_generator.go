package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//ask for number of Exercises
	numExercises := askForExercises()
	//ask for language
	language := askForLanguage()
	//create quiz folder with exercises and tests
	makeExercises(numExercises, language)
}

type filePaths struct {
	exerciseFile string
	testFile     string
	solnsFile    string
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
			formatPrint("That isn't a valid number!")
		}

		if numExercises < 1 || numExercises > 10 {
			formatPrint("You need to pick an integer between 1 and 10")
		} else {
			break
		}
	}
	return numExercises
}

//use stdinnput to ask for language of Exercises to practice
func askForLanguage() string {
	language := "go" // default to go
	for {            //iterate until language is picked
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter language you want to practice (python, go, c, elixir): ")
		language, err := reader.ReadString('\n')
		if err != nil {
			formatPrint("There was an error reading your input")
		}

		match, _ := regexp.MatchString("python|go|c|elixir", language)
		if match {
			break
		} else {
			formatPrint("You need to pick either python, go, c, or elixir")
			fmt.Printf("You printed %v", language)
		}
	}
	return language
}

func makeExercises(numExercises int, language string) {

	//create directory and files
	filePaths := createFiles(language)

	//get fileSuffix
	fileSuffix := getFileSuffix(language)

	//read exercises
	exercisesPath := fmt.Sprintf("./languages/%v/exercises_itemized", language)
	files, err := ioutil.ReadDir(exercisesPath)
	checkErr(err)

	//grab random indices corresponding to files in files slice
	randIdxs := randomIndices(numExercises, len(files))

	//Open exercise and test files for writing

	exerciseFile, err := os.OpenFile(filePaths.exerciseFile, os.O_RDWR, 0666)
	checkErr(err)
	defer exerciseFile.Close()
	testFile, err := os.OpenFile(filePaths.testFile, os.O_RDWR, 0666)
	checkErr(err)
	defer testFile.Close()
	solnFile, err := os.OpenFile(filePaths.solnsFile, os.O_RDWR, 0666)
	checkErr(err)
	defer solnFile.Close()

	initFiles(language, exerciseFile, testFile, solnFile)

	// add exercises and tests to file
	for _, idx := range randIdxs {
		fileName := files[idx].Name()

		//read exercise file
		exercise, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_itemized/%v", language, fileName))
		checkErr(err)

		// read exercise test file
		split := strings.Split(fileName, ".")
		testName := fmt.Sprintf("%v_test%v", split[0], fileSuffix)
		test, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_test/%v", language, testName))
		checkErr(err)

		//read exercise soln file
		solnName := fmt.Sprintf("%v_soln%v", split[0], fileSuffix)
		soln, err := ioutil.ReadFile(fmt.Sprintf("./languages/%v/exercises_solutions/%v", language, solnName))

		//check to make sure files are being read correctly
		fmt.Println("Exercises and Tests for", fileName)

		_, err = exerciseFile.Write(exercise)
		checkErr(err)
		_, err = exerciseFile.Write([]byte("\n \n"))
		checkErr(err)
		_, err = testFile.Write(test)
		checkErr(err)
		_, err = testFile.Write([]byte("\n \n"))
		checkErr(err)
		_, err = solnFile.Write(soln)
		checkErr(err)
		_, err = solnFile.Write([]byte("\n \n"))
		checkErr(err)
	}

	switch language {
	case "c":
		finishTestFileC(testFile)
	case "elixir":
		finishTestFileElixir(testFile)
	}

	createMakeFile()
}

func createFiles(suffix string) *filePaths {
	os.Mkdir("./do_some_exercises", os.ModePerm)
	execsFile := fmt.Sprintf("./do_some_exercises/exercises.%v", suffix)
	execs, err := os.Create(execsFile)
	checkErr(err)
	execs.Close()
	testFile := fmt.Sprintf("./do_some_exercises/exercises_test.%v", suffix)
	tests, err := os.Create(testFile)
	checkErr(err)
	tests.Close()
	solnsFile := fmt.Sprintf("./do_some_exercises/exercises_solns.%v", suffix)
	solns, err := os.Create(solnsFile)
	checkErr(err)
	solns.Close()

	return &filePaths{
		exerciseFile: execsFile,
		testFile:     testFile,
		solnsFile:    solnsFile,
	}
}

func getFileSuffix(language string) string {
	switch language {
	case "python":
		return ".py"
	case "go":
		return ".go"
	case "c":
		return ".c"
	case "elixir":
		return ".ex"
	}
	//default to go
	return ".go"
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
	case "elixir":
		initFilesElixir(exerciseFile, testFile, solnFile)
	}
}

//steps to initialize files differ by language
func initFilesGo(exerciseFile, testFile, solnFile *os.File) {
	//init files with package main
	_, err := exerciseFile.Write([]byte("package goExercises\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("package goExercises\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("import (\n\"fmt\"\n\"testing\"\n\"github.com/google/go-cmp/cmp\"\n)\n\n"))
	checkErr(err)
	_, err = solnFile.Write([]byte("package goExercises\n\n"))
	checkErr(err)
}

func initFilesPython(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("import unittest\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("from exercises import *\n\n"))
	checkErr(err)
}

func initFilesC(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("#include \"exercises.c\"\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("#include <CUnit/CUnit.h>\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("#include <CUnit/Basic.h>\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("int init_suite(void) { return 0; }\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("int clean_suite(void) { return 0; }\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("/************* Your Code goes here **************/\n\n"))
	checkErr(err)
}

func initFilesElixir(exerciseFile, testFile, solnFile *os.File) {
	_, err := testFile.Write([]byte("#Start ExUnit the testrunner\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("ExUnit.start()\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("#create a new test module (test case) and use ExUnit.Case\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("defmodule TestExercises do\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tuse ExUnit.Case, async: true\n\n"))
	checkErr(err)
}

func finishTestFileC(testFile *os.File) {
	_, err := testFile.Write([]byte("/************* Test Runner Code goes here **************/\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("int main ( void )\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("{\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tCU_pSuite pSuite = NULL;\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t/* initialize the CUnit test registry */ \n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tif ( CUE_SUCCESS != CU_initialize_registry() ) \n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t\treturn CU_get_error();\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t/* add a suite to the registry */\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tpSuite = CU_add_suite( \"exercises_test_suite\", init_suite, clean_suite ); \n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tif ( NULL == pSuite ) { \n\t\tCU_cleanup_registry();\n\t\treturn CU_get_error();\n\t}\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t/* add the tests to the suite */\n\tif ( (NULL == CU_add_test(pSuite, \"test_fib\", test_fib)) /* || \n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t\t(NULL == CU_add_test(pSuite, \"max_test_2\", max_test_2)) || \n\t\t(NULL == CU_add_test(pSuite, \"max_test_3\", max_test_3)) */ \n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t)\n{\n\tCU_cleanup_registry(); \n\treturn CU_get_error();\n}\n\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\t// Run all tests using the basic interface\n/tCU_basic_set_mode(CU_BRM_VERBOSE);\n/tCU_basic_run_tests();\n/tprintf(\"\n\");\n/tCU_basic_show_failures(CU_get_failure_list());\n"))
	checkErr(err)
	_, err = testFile.Write([]byte("\tprintf(\"\n\n\"); \n\n\t/* Clean up registry and return */\n\tCU_cleanup_registry();\n/treturn CU_get_error();\n\n}"))
	checkErr(err)

}

func finishTestFileElixir(testFile *os.File) {
	_, err := testFile.Write([]byte("end"))
	checkErr(err)
}

func createMakeFile() {
	makeFilePath := fmt.Sprintf("./do_some_exercises/Makefile")
	makeFile, err := os.Create(makeFilePath)
	checkErr(err)
	defer makeFile.Close()

	_, err = makeFile.Write([]byte("go:\n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("\tgo test \n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("python:\n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("\tpython exercises_test.py \n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("c:\n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("\tgcc exercises_test.c -lcunit -o test ./test \n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("elixir:\n"))
	checkErr(err)
	_, err = makeFile.Write([]byte("\telixir -r exercises.ex exercises_test.exs\n"))
	checkErr(err)
}

// print funcs for formatting
func printSpace() {
	fmt.Printf("\n")
}

func printBorder() {
	fmt.Printf("==================================================================== \n")
}

func formatPrint(printThis string) {
	printBorder()
	fmt.Println(printThis)
	printBorder()
	printSpace()
}

//
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func writeCPreamble() {

}
