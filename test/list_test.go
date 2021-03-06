package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	xs := EmptyStringList().Append("a").ToSlice()
	assert.Equal(t, []string{"a"}, xs)
}

func TestAppendAll(t *testing.T) {
	xs := EmptyStringList().AppendAll("a", "b", "c").ToSlice()
	assert.Equal(t, []string{"a", "b", "c"}, xs)
}

func TestAppendSlice(t *testing.T) {
	xs := EmptyStringList().AppendSlice([]string{"a", "b", "c"}).ToSlice()
	assert.Equal(t, []string{"a", "b", "c"}, xs)
}

func TestIntersperse(t *testing.T) {
	xs := EmptyStringList().AppendAll("a", "b", "c").Intersperse(",").ToSlice()
	assert.Equal(t, []string{"a", ",", "b", ",", "c"}, xs)
}

func TestFill(t *testing.T) {
	xs := EmptyStringList().Fill(3, "a").ToSlice()
	assert.Equal(t, []string{"a", "a", "a"}, xs)
}

func TestTabulate(t *testing.T) {
	xs := EmptyIntList().Tablulate(5, func(x int) int { return x * x }).ToSlice()
	ys := EmptyStringList().Tablulate(5, func(x int) string { return fmt.Sprintf("%v", x*x) }).ToSlice()
	assert.Equal(t, []int{0, 1, 4, 9, 16}, xs)
	assert.Equal(t, []string{"0", "1", "4", "9", "16"}, ys)
}

func TestPartition(t *testing.T) {
	xs := EmptyIntList().Concat(1, 2, 3, 4, 5, 6, 7)
	ys, zs := xs.Partition(func(x int) bool { return x > 3 })
	assert.Equal(t, []int{4, 5, 6, 7}, ys.ToSlice())
	assert.Equal(t, []int{1, 2, 3}, zs.ToSlice())
}

func TestMkString(t *testing.T) {
	xs := EmptyIntList().Concat(1, 2, 3, 4)
	str := xs.MkString(func(i int) string {
		return fmt.Sprintf("%v", i)
	}).Str()
	assert.Equal(t, "1234", str)
}

func TestDistinct(t *testing.T) {
	xs := EmptyIntList().Concat(1, 2, 3, 1, 2, 3, 4, 1, 2, 3).Distinct().ToSlice()
	assert.Equal(t, []int{1, 2, 3, 4}, xs)
}

func TestIntersect(t *testing.T) {
	xs := EmptyIntList().Concat(1, 2, 4, 5)
	ys := EmptyIntList().Concat(2, 3, 4)
	zs := xs.Intersect(ys).ToSlice()
	assert.Equal(t, []int{2, 4}, zs)
}

func TestSlice(t *testing.T) {
	xs := EmptyIntList().Concat(1, 2, 3, 4, 5).Slice(1, 2).ToSlice()
	assert.Equal(t, []int{2, 3}, xs)
}

func TestHeadOption(t *testing.T) {
	xs := EmptyIntList().AppendAll(1, 2, 3)
	assert.Equal(t, IntSome{1}, xs.HeadOption())
	ys := EmptyIntList()
	assert.Equal(t, noneInt, ys.HeadOption())
}

func TestFlatMap(t *testing.T) {
	xs := EmptyIntList().
		AppendAll(1, 2, 3).
		FlatMapToIntList(func(x int) IntList {
			return EmptyIntList().AppendAll(1, 2, 3)
		})
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 1, 2, 3}, xs.ToSlice())

	ys := EmptyIntList().
		AppendAll(1, 2, 3).
		FlatMapToStringList(func(x int) StringList {
			return EmptyStringList().AppendAll("a", "b", "c")
		})
	assert.Equal(t, []string{"a", "b", "c", "a", "b", "c", "a", "b", "c"}, ys.ToSlice())
}

func TestFlatMapToPerson(t *testing.T) {
	xs := EmptyStringList().
		AppendAll("foo", "bar").
		FlatMapToCatList(func(s string) CatList {
			return EmptyCatList().
				Append(Cat{Name: s}).
				Append(Cat{Name: s})
		})
	assert.Equal(t, []Cat{{Name: "foo"}, {Name: "foo"}, {Name: "bar"}, {Name: "bar"}}, xs.ToSlice())
}

func TestJoin(t *testing.T) {
	x := String("a/b/c").
		Split("/").
		Init().
		Join("/").
		Append("/").
		Str()
	assert.Equal(t, "a/b/", x)
}

func TestSort(t *testing.T) {
	xs := EmptyCatList().AppendSlice([]Cat{
		{Name: "a", Age: 1},
		{Name: "d", Age: 4},
		{Name: "c", Age: 3},
		{Name: "b", Age: 2},
	})

	xs.Sort(func(i Cat, j Cat) bool {
		return i.Age < j.Age
	})

	assert.Equal(t, []Cat{{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}}, xs.ToSlice())
}

var result IntList

func BenchmarkMapStringList(b *testing.B) {
	var xs IntList
	for n := 0; n < b.N; n++ {
		xs = EmptyStringList().
			AppendAll("a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh").
			MapToInt(func(s string) int {
				return len(s)
			})
	}
	result = xs
}

var result2 []int

func BenchmarkRangeString(b *testing.B) {
	var ys []int
	for n := 0; n < b.N; n++ {
		xs := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh"}
		ys = make([]int, 0)
		for _, x := range xs {
			ys = append(ys, len(x))
		}

	}
	result2 = ys
}
