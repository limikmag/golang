package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	cases := []struct {
		name          string
		array         []int
		expectedArray []int
		expLen        int
	}{
		{"Basic", []int{0, 0, 1, 1, 1, 2, 5}, []int{0, 1, 2, 5}, 4},
	}

	for _, tc := range cases {
		resAr, _ := removeDuplicates(tc.array)
		assert.Equal(t, tc.expectedArray, resAr, "returned %v but expected %v", resAr, tc.expectedArray)

	}
}
