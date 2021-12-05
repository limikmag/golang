package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Binary(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 67, 122, 888, 999, 1000, 10002, 11110}

	index := Binary(sortedSlice, 122, 0, len(sortedSlice)-1)
	assert.Equal(index, 4, fmt.Sprintf("index should be 4 but returned %v", index))
}
