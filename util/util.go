package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	case "ocaml":
		return ".ml"
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

func RemoveOldExercises() error {
	dir := "./do_some_exercises"
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
