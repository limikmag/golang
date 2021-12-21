package sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {

	cases := []struct {
		name     string
		sample   []int
		sum      int
		expected bool
	}{
		{"TestWithEmptyTable", []int{}, 10, false},
		{"TestPositive", []int{1, 4, 5, 7, 1, 2, 3, 7}, 8, true},
		{"TestNegative", []int{1, 4, 5, 7, 1, 2, 3, 7}, 122, false},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, Sum(tc.sample, tc.sum), fmt.Sprintf("For %v and %v we should return %v but returned %v", tc.sample, tc.sum, tc.expected, Sum(tc.sample, tc.sum)))
	}
}
