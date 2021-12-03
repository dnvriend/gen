// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"strings"
)

type IntList []int

func EmptyIntList() IntList {
	return make([]int, 0)
}

// append an element to the end of the list
func (rcv IntList) Append(x int) IntList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv IntList) AppendAll(xs ...int) IntList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv IntList) Concat(xs ...int) IntList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv IntList) AppendSlice(xs []int) IntList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv IntList) ConcatSlice(xs []int) IntList {
	return rcv.AppendSlice(xs)
}

func (rcv IntList) Reverse() IntList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv IntList) Head() int {
	return rcv[0]
}

func (rcv IntList) HeadOption() IntOption {
	if len(rcv) == 0 {
		return OptionOfInt(nil)
	}
	return OptionOfInt(&rcv[0])
}

func (rcv IntList) Last() int {
	return rcv[len(rcv)-1]
}

// returns the initial part of the collection, without the last element
func (rcv IntList) Init() IntList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv IntList) Tail() IntList {
	return rcv[1:]
}

// Selects all elements of this list which satisfy a predicate.
func (rcv IntList) Filter(fn func(int) bool) IntList {
	ys := EmptyIntList()
	rcv.ForEach(func(v int) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv IntList) TakeWhile(fn func(int) bool) IntList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv IntList) FilterNot(fn func(int) bool) IntList {
	return rcv.Filter(func(x int) bool { return !fn(x) })
}

// alias for FilterNot
func (rcv IntList) DropWhile(fn func(int) bool) IntList {
	return rcv.FilterNot(fn)
}

func (rcv IntList) ForEach(fn func(int)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv IntList) ForEachWithIndex(fn func(int, int)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv IntList) ForEachWithLastFlag(fn func(bool, int)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv IntList) ForEachP(fn func(int)) {
	rcv.ForEachPP(10, fn)
}

func (rcv IntList) ForEachPP(parallelism int, fn func(int)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv IntList) ForEachPPP(parallelism int, fn func(int), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan int, nrJobs)
	output := make(chan bool, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				fn(x)
				output <- true
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x int) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv IntList) Count() int {
	return len(rcv)
}

func (rcv IntList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv IntList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv IntList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of int
func (rcv IntList) ToSlice() []int {
	return rcv
}

// returns the element at the index
func (rcv IntList) Apply(x int) int {
	return rcv[x]
}

func (rcv IntList) ApplyOrElse(x int, fn func() int) int {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv IntList) FoldLeft(zero int, fn func(acc int, x int) int) int {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv IntList) FoldRight(zero int, fn func(acc int, x int) int) int {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv IntList) Intersperse(a int) IntList {
	ys := EmptyIntList()
	rcv.ForEachWithLastFlag(func(last bool, x int) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv IntList) Fill(num int, a int) IntList {
	xs := EmptyIntList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv IntList) Tablulate(num int, fn func(int) int) IntList {
	xs := EmptyIntList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv IntList) Partition(fn func(int) bool) (IntList, IntList) {
	xs := EmptyIntList()
	ys := EmptyIntList()
	rcv.ForEach(func(x int) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv IntList) MkString(fn func(int) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x int) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv IntList) RangeOf(from int, to int, fn func(int) int) IntList {
	xs := EmptyIntList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv IntList) Contains(a int) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv IntList) ContainsNot(a int) bool {
	return !rcv.Contains(a)
}

func (rcv IntList) Distinct() IntList {
	xs := EmptyIntList()
	rcv.ForEach(func(x int) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv IntList) Intersect(xs IntList) IntList {
	ys := EmptyIntList()
	rcv.ForEach(func(x int) {
		xs.ForEach(func(y int) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv IntList) Slice(from int, to int) IntList {
	return rcv[from : to+1]
}

func (rcv IntList) FlatMapToIntList(fn func(int) IntList) IntList {
	xs := EmptyIntList()
	rcv.ForEach(func(x int) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv IntList) MapToInt(fn func(int) int) IntList {
	xs := EmptyIntList()
	rcv.ForEach(func(x int) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv IntList) MapToIntWithIndex(fn func(int, int) int) IntList {
	xs := EmptyIntList()
	rcv.ForEachWithIndex(func(i int, x int) {
		xs = xs.Append(fn(i, x))
	})
	return xs
}

func (rcv IntList) MapToIntP(mapFn func(int) int) IntList {
	return rcv.MapToIntPP(10, mapFn)
}

func (rcv IntList) MapToIntPP(parallelism int, mapFn func(int) int) IntList {
	return rcv.MapToIntPPP(parallelism, mapFn, func() {})
}

func (rcv IntList) MapToIntPPP(parallelism int, mapFn func(int) int, progressFn func()) IntList {
	nrJobs := rcv.Count()
	input := make(chan int, nrJobs)
	output := make(chan int, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x int) {
		input <- x
	})
	close(input)

	xs := EmptyIntList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

// implementation of 'sort.Interface'
func (rcv IntList) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv IntList) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var IntListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv IntList) Less(i, j int) bool {
	return IntListLessFunc(i, j)
}

// i and j are two objects that need to be compared,
// and based on that comparison the List will be sorted
func (rcv IntList) Sort(fn func(i int, j int) bool) IntList {
	IntListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}

func (rcv IntList) MapToString(fn func(int) string) StringList {
	ys := make([]string, 0)
	for _, x := range rcv {
		ys = append(ys, fn(x))
	}
	return ys
}

func (rcv IntList) MapToStringWithIndex(fn func(int, int) string) StringList {
	ys := make([]string, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i, x))
	}
	return ys
}

func (rcv IntList) MapToStringWithLastFlag(fn func(bool, int) string) StringList {
	ys := make([]string, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i+1 == len(rcv), x))
	}
	return ys
}

func (rcv IntList) MapToStringP(mapFn func(int) string) StringList {
	return rcv.MapToStringPP(10, mapFn)
}

func (rcv IntList) MapToStringPP(parallelism int, mapFn func(int) string) StringList {
	return rcv.MapToStringPPP(parallelism, mapFn, func() {})
}

func (rcv IntList) MapToStringPPP(parallelism int, mapFn func(int) string, progressFn func()) StringList {
	nrJobs := rcv.Count()
	input := make(chan int, nrJobs)
	output := make(chan string, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x int) {
		input <- x
	})
	close(input)

	xs := EmptyStringList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

func (rcv IntList) FlatMapToStringList(fn func(int) StringList) StringList {
	xs := EmptyStringList()
	rcv.ForEach(func(x int) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv IntList) FlatMapToCatList(fn func(int) CatList) CatList {
	xs := EmptyCatList()
	rcv.ForEach(func(x int) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv IntList) Range(from int, to int) IntList {
	xs := EmptyIntList()
	for i := from; i <= to; i++ {
		xs = xs.Append(i)
	}
	return xs
}
