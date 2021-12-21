package shortestsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestSubstring(t *testing.T) {
	cases := []struct {
		name          string
		text          string
		words         []string
		caseSensitive bool
		expected      string
	}{
		{"TestWithNotPresentWordCaseSensitive", "abra kadabra 2 kije, kto sie obudzi ten ginie", []string{"abra", "michal"}, true, ""},
		{"TestWithNotPresentWordCaseNonSensitive", "abra kadabra 2 kije, kto sie obudzi ten ginie", []string{"abra", "michal"}, false, ""},
		{"TestWithPresentCaseSensitive", "abra kadabra 2 kije, michal kmak michal", []string{"kadabra", "kije"}, true, "kadabra 2 kije"},
		{"TestWithPresentCaseSensitive", "abra kadabra 2 kije, michal kmak michal", []string{"kadabra", "kije", "michal"}, true, "kadabra 2 kije, michal"},
		{"TestWithPresentCaseSensitive", "abra kadabra 2 kije, michal kmak michal ala", []string{"kadabra", "kije", "michal", "ala"}, true, "kadabra 2 kije, michal kmak michal ala"},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, shortestSubstring(tc.text, tc.words, tc.caseSensitive), fmt.Sprintf("For %v and %v we should return %v but returned %v", tc.text, tc.words, tc.expected, shortestSubstring(tc.text, tc.words, tc.caseSensitive)))
	}
}

func Test_allOccurrencess(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		word     string
		expected []int
	}{
		{"TestEmptyText", "", "michal", []int{}},
		{"TestEmptyText", "michal michal kwora michal", "michal", []int{0, 7, 20}},
		{"TestEmptyWord", "michal michal kwora michal", "", []int{}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, allOccurrencess(tc.text, tc.word), fmt.Sprintf("For %v and %v we should return %v but returned %v", tc.text, tc.word, tc.expected, allOccurrencess(tc.text, tc.word)))
	}
}
