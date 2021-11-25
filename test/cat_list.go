// Generated code; DO NOT EDIT.
package test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
)

type CatList []Cat

func EmptyCatList() CatList {
	return make([]Cat, 0)
}

// append an element to the end of the list
func (rcv CatList) Append(x Cat) CatList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv CatList) AppendAll(xs ...Cat) CatList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv CatList) Concat(xs ...Cat) CatList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv CatList) AppendSlice(xs []Cat) CatList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv CatList) ConcatSlice(xs []Cat) CatList {
	return rcv.AppendSlice(xs)
}

func (rcv CatList) Reverse() CatList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv CatList) Head() Cat {
	return rcv[0]
}

func (rcv CatList) HeadOption() CatOption {
	if len(rcv) == 0 {
		return OptionOfCat(nil)
	}
	return OptionOfCat(&rcv[0])
}

func (rcv CatList) Last() Cat {
	return rcv[len(rcv)-1]
}

// returns the initial part of the collection, without the last element
func (rcv CatList) Init() CatList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv CatList) Tail() CatList {
	return rcv[1:]
}

// Selects all elements of this list which satisfy a predicate.
func (rcv CatList) Filter(fn func(Cat) bool) CatList {
	ys := EmptyCatList()
	rcv.ForEach(func(v Cat) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv CatList) TakeWhile(fn func(Cat) bool) CatList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv CatList) FilterNot(fn func(Cat) bool) CatList {
	return rcv.Filter(func(x Cat) bool { return !fn(x) })
}

// alias for FilterNot
func (rcv CatList) DropWhile(fn func(Cat) bool) CatList {
	return rcv.FilterNot(fn)
}

func (rcv CatList) ForEach(fn func(Cat)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv CatList) ForEachWithIndex(fn func(int, Cat)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv CatList) ForEachWithLastFlag(fn func(bool, Cat)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv CatList) ForEachP(fn func(Cat)) {
	rcv.ForEachPP(10, fn)
}

func (rcv CatList) ForEachPP(parallelism int, fn func(Cat)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv CatList) ForEachPPP(parallelism int, fn func(Cat), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan Cat, nrJobs)
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
	rcv.ForEach(func(x Cat) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv CatList) Count() int {
	return len(rcv)
}

func (rcv CatList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv CatList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv CatList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of Cat
func (rcv CatList) ToSlice() []Cat {
	return rcv
}

// returns the element at the index
func (rcv CatList) Apply(x int) Cat {
	return rcv[x]
}

func (rcv CatList) ApplyOrElse(x int, fn func() Cat) Cat {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv CatList) FoldLeft(zero Cat, fn func(acc Cat, x Cat) Cat) Cat {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv CatList) FoldRight(zero Cat, fn func(acc Cat, x Cat) Cat) Cat {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv CatList) Intersperse(a Cat) CatList {
	ys := EmptyCatList()
	rcv.ForEachWithLastFlag(func(last bool, x Cat) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv CatList) Fill(num int, a Cat) CatList {
	xs := EmptyCatList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv CatList) Tablulate(num int, fn func(int) Cat) CatList {
	xs := EmptyCatList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv CatList) Partition(fn func(Cat) bool) (CatList, CatList) {
	xs := EmptyCatList()
	ys := EmptyCatList()
	rcv.ForEach(func(x Cat) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv CatList) MkString() String {
	var builder strings.Builder
	rcv.ForEach(func(x Cat) {
		builder.WriteString(fmt.Sprintf("%v", x))
	})
	return String(builder.String())
}

func (rcv CatList) RangeOf(from int, to int, fn func(int) Cat) CatList {
	xs := EmptyCatList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv CatList) Contains(a Cat) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv CatList) ContainsNot(a Cat) bool {
	return !rcv.Contains(a)
}

func (rcv CatList) Distinct() CatList {
	xs := EmptyCatList()
	rcv.ForEach(func(x Cat) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv CatList) Intersect(xs CatList) CatList {
	ys := EmptyCatList()
	rcv.ForEach(func(x Cat) {
		xs.ForEach(func(y Cat) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv CatList) Slice(from int, to int) CatList {
	return rcv[from : to+1]
}

func (rcv CatList) FlatMapToCatList(fn func(Cat) CatList) CatList {
	xs := EmptyCatList()
	rcv.ForEach(func(x Cat) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv CatList) MapToCat(fn func(Cat) Cat) CatList {
	xs := EmptyCatList()
	rcv.ForEach(func(x Cat) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv CatList) MapToCatP(mapFn func(Cat) Cat) CatList {
	return rcv.MapToCatPP(10, mapFn)
}

func (rcv CatList) MapToCatPP(parallelism int, mapFn func(Cat) Cat) CatList {
	return rcv.MapToCatPPP(parallelism, mapFn, func() {})
}

func (rcv CatList) MapToCatPPP(parallelism int, mapFn func(Cat) Cat, progressFn func()) CatList {
	nrJobs := rcv.Count()
	input := make(chan Cat, nrJobs)
	output := make(chan Cat, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x Cat) {
		input <- x
	})
	close(input)

	xs := EmptyCatList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}
