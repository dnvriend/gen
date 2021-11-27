package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeList(t *testing.T) {
	xs := EmptyIntList().Range(1, 4).ToSlice()
	ys := EmptyIntList().RangeOf(1, 4, func(x int) int { return x }).ToSlice()
	assert.Equal(t, []int{1, 2, 3, 4}, xs)
	assert.Equal(t, []int{1, 2, 3, 4}, ys)
}

func TestRange(t *testing.T) {
	xs := Range(1, 5).ToIntList().ToSlice()
	assert.Equal(t, []int{1, 2, 3, 4}, xs)
}

func TestRangeWithStep(t *testing.T) {
	tests := []struct {
		input    int
		from     int
		to       int
		expected []int
		message  string
	}{
		{1, 1, 5, []int{1, 2, 3, 4}, "step one"},
		{2, 1, 5, []int{1, 3}, "step two"},
		{3, 1, 10, []int{1, 4, 7}, "step three"},
		{4, 1, 25, []int{1, 5, 9, 13, 17, 21}, "step four"},
		{5, 1, 30, []int{1, 6, 11, 16, 21, 26}, "step five"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Range(test.from, test.to).WithStep(test.input).ToIntList().ToSlice(), test.message)
	}
}
