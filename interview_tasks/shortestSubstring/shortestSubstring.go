package shortestsubstring

import (
	"strings"
)

func shortestSubstring(text string, words []string, caseSensitive bool) string {

	// non case sensitive
	if !caseSensitive {
		text, words = noCS(text, words)
	}

	// in case when some of the words are not present in text
	if !allWordsPresent(text, words) {
		return ""
	}

	positions := make(map[string][]int)

	for _, word := range words {
		positions[word] = allOccurrencess(text, word)
	}

	startPositions := []int{}
	endPositions := []int{}

	for key, value := range positions {
		startPositions = append(startPositions, value...)
		endPositions = append(endPositions, func(len int, val []int) []int {
			res := []int{}
			for _, num := range val {
				res = append(res, num+len)
			}
			return res
		}(len(key), value)...)
	}

	return minSubsequence(generateRightOptions(text, words, startPositions, endPositions))
}

func noCS(text string, words []string) (string, []string) {
	newText := strings.ToLower(text)
	newWords := []string{}
	for _, word := range words {
		newWords = append(newWords, strings.ToLower(word))
	}

	return newText, newWords

}

func allWordsPresent(text string, words []string) bool {
	for _, word := range words {
		if !strings.Contains(text, word) {
			return false
		}
	}

	return true
}

func allOccurrencess(text string, word string) []int {
	res := []int{}

	if word == "" {
		return res
	}

	occurence := 0
	for {
		index := strings.Index(text, word)
		if index == -1 {
			break
		}
		if occurence > 0 {
			res = append(res, index+res[occurence-1]+len(word))
		} else {
			res = append(res, index)
		}

		text = text[index+len(word):]
		occurence++
	}
	return res
}

func generateRightOptions(text string, words []string, start []int, end []int) []string {
	res := []string{}
	for _, i := range start {
		for _, j := range end {
			if i < j && allWordsPresent(text[i:j], words) {
				res = append(res, text[i:j])
			}
		}
	}

	return res
}

func minSubsequence(subs []string) string {
	min := int(10e10)
	smallestSub := ""
	for _, sub := range subs {
		if len(sub) < min {
			min = len(sub)
			smallestSub = sub
		}
	}

	return smallestSub
}
