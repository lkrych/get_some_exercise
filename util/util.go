package util

import (
	"fmt"
	"log"
)

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
