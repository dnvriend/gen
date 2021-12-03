// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"strings"
)

type FloatList []Float

func EmptyFloatList() FloatList {
	return make([]Float, 0)
}

// append an element to the end of the list
func (rcv FloatList) Append(x Float) FloatList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv FloatList) AppendAll(xs ...Float) FloatList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv FloatList) Concat(xs ...Float) FloatList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv FloatList) AppendSlice(xs []Float) FloatList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv FloatList) ConcatSlice(xs []Float) FloatList {
	return rcv.AppendSlice(xs)
}

func (rcv FloatList) Reverse() FloatList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv FloatList) Head() Float {
	return rcv[0]
}

func (rcv FloatList) HeadOption() FloatOption {
	if len(rcv) == 0 {
		return OptionOfFloat(nil)
	}
	return OptionOfFloat(&rcv[0])
}

func (rcv FloatList) Last() Float {
	return rcv[len(rcv)-1]
}

// returns the initial part of the collection, without the last element
func (rcv FloatList) Init() FloatList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv FloatList) Tail() FloatList {
	return rcv[1:]
}

// Selects all elements of this list which satisfy a predicate.
func (rcv FloatList) Filter(fn func(Float) bool) FloatList {
	ys := EmptyFloatList()
	rcv.ForEach(func(v Float) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv FloatList) TakeWhile(fn func(Float) bool) FloatList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv FloatList) FilterNot(fn func(Float) bool) FloatList {
	return rcv.Filter(func(x Float) bool { return !fn(x) })
}

// alias for FilterNot
func (rcv FloatList) DropWhile(fn func(Float) bool) FloatList {
	return rcv.FilterNot(fn)
}

func (rcv FloatList) ForEach(fn func(Float)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv FloatList) ForEachWithIndex(fn func(int, Float)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv FloatList) ForEachWithLastFlag(fn func(bool, Float)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv FloatList) ForEachP(fn func(Float)) {
	rcv.ForEachPP(10, fn)
}

func (rcv FloatList) ForEachPP(parallelism int, fn func(Float)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv FloatList) ForEachPPP(parallelism int, fn func(Float), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan Float, nrJobs)
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
	rcv.ForEach(func(x Float) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv FloatList) Count() int {
	return len(rcv)
}

func (rcv FloatList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv FloatList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv FloatList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of Float
func (rcv FloatList) ToSlice() []Float {
	return rcv
}

// returns the element at the index
func (rcv FloatList) Apply(x int) Float {
	return rcv[x]
}

func (rcv FloatList) ApplyOrElse(x int, fn func() Float) Float {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv FloatList) FoldLeft(zero Float, fn func(acc Float, x Float) Float) Float {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv FloatList) FoldRight(zero Float, fn func(acc Float, x Float) Float) Float {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv FloatList) Intersperse(a Float) FloatList {
	ys := EmptyFloatList()
	rcv.ForEachWithLastFlag(func(last bool, x Float) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv FloatList) Fill(num int, a Float) FloatList {
	xs := EmptyFloatList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv FloatList) Tablulate(num int, fn func(int) Float) FloatList {
	xs := EmptyFloatList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv FloatList) Partition(fn func(Float) bool) (FloatList, FloatList) {
	xs := EmptyFloatList()
	ys := EmptyFloatList()
	rcv.ForEach(func(x Float) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv FloatList) MkString(fn func(Float) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x Float) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv FloatList) RangeOf(from int, to int, fn func(int) Float) FloatList {
	xs := EmptyFloatList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv FloatList) Contains(a Float) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv FloatList) ContainsNot(a Float) bool {
	return !rcv.Contains(a)
}

func (rcv FloatList) Distinct() FloatList {
	xs := EmptyFloatList()
	rcv.ForEach(func(x Float) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv FloatList) Intersect(xs FloatList) FloatList {
	ys := EmptyFloatList()
	rcv.ForEach(func(x Float) {
		xs.ForEach(func(y Float) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv FloatList) Slice(from int, to int) FloatList {
	return rcv[from : to+1]
}

func (rcv FloatList) FlatMapToFloatList(fn func(Float) FloatList) FloatList {
	xs := EmptyFloatList()
	rcv.ForEach(func(x Float) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv FloatList) MapToFloat(fn func(Float) Float) FloatList {
	xs := EmptyFloatList()
	rcv.ForEach(func(x Float) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv FloatList) MapToFloatWithIndex(fn func(int, Float) Float) FloatList {
	xs := EmptyFloatList()
	rcv.ForEachWithIndex(func(i int, x Float) {
		xs = xs.Append(fn(i, x))
	})
	return xs
}

func (rcv FloatList) MapToFloatP(mapFn func(Float) Float) FloatList {
	return rcv.MapToFloatPP(10, mapFn)
}

func (rcv FloatList) MapToFloatPP(parallelism int, mapFn func(Float) Float) FloatList {
	return rcv.MapToFloatPPP(parallelism, mapFn, func() {})
}

func (rcv FloatList) MapToFloatPPP(parallelism int, mapFn func(Float) Float, progressFn func()) FloatList {
	nrJobs := rcv.Count()
	input := make(chan Float, nrJobs)
	output := make(chan Float, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x Float) {
		input <- x
	})
	close(input)

	xs := EmptyFloatList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

// implementation of 'sort.Interface'
func (rcv FloatList) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv FloatList) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var FloatListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv FloatList) Less(i, j int) bool {
	return FloatListLessFunc(i, j)
}

// i and j are two objects that need to be compared,
// and based on that comparison the List will be sorted
func (rcv FloatList) Sort(fn func(i Float, j Float) bool) FloatList {
	FloatListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}
