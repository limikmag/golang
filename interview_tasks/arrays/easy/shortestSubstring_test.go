package easy

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
		{
			"TestBig",
			`Litwo, Ojczyzno moja! ty jesteś jak zdrowie;
Ile cię trzeba cenić, ten tylko się dowie,
Kto cię stracił. Dziś piękność twą w całej ozdobie
Widzę i opisuję, bo tęsknię po tobie.

Panno święta, co Jasnej bronisz Częstochowy
I w Ostrej świecisz Bramie! Ty, co gród zamkowy
Nowogródzki ochraniasz z jego wiernym ludem!
Jak mnie dziecko do zdrowia powróciłaś cudem
(— Gdy od płaczącej matki, pod Twoją opiekę
Ofiarowany martwą podniosłem powiekę;
I zaraz mogłem pieszo, do Twych świątyń progu
Iść za wrócone życie podziękować Bogu —)
Tak nas powrócisz cudem na Ojczyzny łono!...
Tymczasem, przenoś moją duszę utęsknioną
Do tych pagórków leśnych, do tych łąk zielonych,
Szeroko nad błękitnym Niemnem rozciągnionych;
Do tych pól malowanych zbożem rozmaitem,
Wyzłacanych pszenicą, posrebrzanych żytem;
Gdzie bursztynowy świerzop, gryka jak śnieg biała,
Gdzie panieńskim rumieńcem dzięcielina pała,
A wszystko przepasane jakby wstęgą, miedzą
Zieloną, na niej zrzadka ciche grusze siedzą.`,
			[]string{"cudem", "świerzop"},
			true,
			`cudem na Ojczyzny łono!...
Tymczasem, przenoś moją duszę utęsknioną
Do tych pagórków leśnych, do tych łąk zielonych,
Szeroko nad błękitnym Niemnem rozciągnionych;
Do tych pól malowanych zbożem rozmaitem,
Wyzłacanych pszenicą, posrebrzanych żytem;
Gdzie bursztynowy świerzop`,
		},
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
