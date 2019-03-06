package util

import (
	"fmt"
	"log"
)

func GetFileSuffix(language string) string {
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

// print funcs for formatting
func PrintSpace() {
	fmt.Printf("\n")
}

func printBorder() {
	fmt.Printf("==================================================================== \n")
}

func FormatPrint(printThis string) {
	printBorder()
	fmt.Println(printThis)
	printBorder()
	PrintSpace()
}

//
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
