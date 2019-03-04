// import (
// 	"io/ioutil"
// 	"log"
// 	"strconv"
// 	"strings"
// )

// func checkSum(multiArr [][]int) int {
// 	checksum := 0
// 	for idx, arr := range multiArr {
// 		max := -100000
// 		min := 100000
// 		for i := 0; i < len(multiArr[idx]); i++ {
// 			if arr[i] > max {
// 				max = arr[i]
// 			}
// 			if arr[i] < min {
// 				min = arr[i]
// 			}
// 		}
// 		checksum += (max - min)
// 	}
// 	return checksum
// }

// func checkSumAdvent() [][]int {
// 	content, err := ioutil.ReadFile("./advent_input/checksum_input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	arrays := strings.Split(string(content), "\n")
// 	multiarr := make([][]int, len(arrays)-1)

// 	for i := 0; i < len(arrays)-1; i++ {
// 		split := strings.Split(string(arrays[i]), "	")
// 		intSplit := make([]int, len(split))
// 		for j := 0; j < len(split); j++ {
// 			intSplit[j], _ = strconv.Atoi(split[j])
// 		}
// 		multiarr[i] = intSplit
// 	}
// 	return multiarr
// }