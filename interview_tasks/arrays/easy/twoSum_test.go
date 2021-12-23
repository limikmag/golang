package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	cases := []struct {
		name   string
		list   []int
		target int
		result []int
	}{
		{"Basic", []int{1, 2, 3}, 5, []int{1, 2}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.result, twoSum(tc.list, tc.target), "expected %v but result %v", tc.result, twoSum(tc.list, tc.target))
	}
}
