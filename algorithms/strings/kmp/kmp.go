package main

import "fmt"

// User defined.
// Set to true to read input from two command line arguments
// Set to false to read input from two files "pattern.txt" and "text.txt"

// const isTakingInputFromCommandLine bool = true

const notFoundPosition int = -1

type Result struct {
	resultPosition     int
	numberOfComparison int
}

// Table building algorithm.
// Takes word to be analyzed and table to be filled.
func kmpTable(word string) (t []int) {
	t = make([]int, len(word))
	pos, cnd := 2, 0
	t[0], t[1] = -1, 0
	for pos < len(word) {
		if word[pos-1] == word[cnd] {
			cnd++
			t[pos] = cnd
			pos++
		} else if cnd > 0 {
			cnd = t[cnd]
		} else {
			t[pos] = 0
			pos++
		}
	}
	return t
}

func main() {
	fmt.Println(kmpTable("abcd"))
}
