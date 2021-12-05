package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Binary_1(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 67, 122, 888, 999, 1000, 10002, 11110}

	index := Binary(sortedSlice, 122, 0, len(sortedSlice)-1)
	assert.Equal(t, 4, index, fmt.Sprintf("index should be 4 but returned %v", index))
}

func Test_Binary_2(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 67, 122, 888, 999, 1000, 10002, 11110}

	index := Binary(sortedSlice, 123, 0, len(sortedSlice)-1)
	assert.Equal(t, -1, index, fmt.Sprintf("index should be 4 but returned %v", index))
}

func Test_BinaryIterative_1(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 67, 122, 888, 999, 1000, 10002, 11110}

	index := BinaryIterative(sortedSlice, 122, 0, len(sortedSlice)-1)
	assert.Equal(t, 4, index, fmt.Sprintf("index should be 4 but returned %v", index))
}

func Test_BinaryIterative_2(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 67, 122, 888, 999, 1000, 10002, 11110}

	index := BinaryIterative(sortedSlice, 123, 0, len(sortedSlice)-1)
	assert.Equal(t, -1, index, fmt.Sprintf("index should be 4 but returned %v", index))
}
