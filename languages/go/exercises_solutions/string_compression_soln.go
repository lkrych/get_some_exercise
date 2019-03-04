//before building the string, you should check to make sure that you need to build the string
// func stringCompression(in string) out string {
// 	if countCompressionLength(in) > len(in) {
// 		return in
// 	}

// 	compressed := []rune{}
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(string) || in[i] != in[i+1]) {
// 			compressed = append(compressed, in[i], countConsec)
// 			countConsec = 0
// 		}
// 	}
// 	return string(compressed)
// }

// func countCompressionLength(in string) bool {
// 	compressedLength := 0
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(in) || in[i] != in[i+1]) {
// 			compressedLength += (1 + len(string(countConsec)))
// 			countConsec = 0
// 		}
// 	}
// }