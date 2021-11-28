package test

import (
	"math"
	"testing"
)

//go:generate gen float -p test
//go:generate gen list -p test -t Float
//go:generate gen option -p test -t Float
func TestFloat(t *testing.T) {
	EmptyFloatList().RangeOf(1, 100, func(i int) Float {
		x := math.Pi * (float64)(i)
		return ToFloat(x)
	}).
		MapToFloat(func(float Float) Float {
			return float.Cos()
		}).
		ForEach(func(float Float) {
			float.Println()
		})
}
