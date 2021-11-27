// Generated code; DO NOT EDIT.
package main

import (
	"sort"
	"strings"
	"github.com/google/go-cmp/cmp"
	
)

type DadJokeList []DadJoke

func EmptyDadJokeList() DadJokeList {
	return make([]DadJoke, 0)
}

// append an element to the end of the list
func (rcv DadJokeList) Append(x DadJoke) DadJokeList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv DadJokeList) AppendAll(xs ...DadJoke) DadJokeList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv DadJokeList) Concat(xs ...DadJoke) DadJokeList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv DadJokeList) AppendSlice(xs []DadJoke) DadJokeList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv DadJokeList) ConcatSlice(xs []DadJoke) DadJokeList {
	return rcv.AppendSlice(xs)
}

func (rcv DadJokeList) Reverse() DadJokeList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv DadJokeList) Head() DadJoke {
	return rcv[0] 
}

func (rcv DadJokeList) HeadOption() DadJokeOption {
	if len(rcv) == 0 {
		return OptionOfDadJoke(nil)
	} 
	return OptionOfDadJoke(&rcv[0])
}

func (rcv DadJokeList) Last() DadJoke {
	return rcv[len(rcv)-1] 
}

// returns the initial part of the collection, without the last element
func (rcv DadJokeList) Init() DadJokeList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv DadJokeList) Tail() DadJokeList {
	return rcv[1:]
} 

// Selects all elements of this list which satisfy a predicate.
func (rcv DadJokeList) Filter(fn func(DadJoke) bool) DadJokeList {
	ys := EmptyDadJokeList()
 	rcv.ForEach(func(v DadJoke) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv DadJokeList) TakeWhile(fn func(DadJoke) bool) DadJokeList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv DadJokeList) FilterNot(fn func(DadJoke) bool) DadJokeList {
	return rcv.Filter(func (x DadJoke) bool { return !fn(x)})
}

// alias for FilterNot
func (rcv DadJokeList) DropWhile(fn func(DadJoke) bool) DadJokeList {
	return rcv.FilterNot(fn)
}

func (rcv DadJokeList) ForEach(fn func(DadJoke)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv DadJokeList) ForEachWithIndex(fn func(int, DadJoke)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv DadJokeList) ForEachWithLastFlag(fn func(bool, DadJoke)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv DadJokeList) ForEachP(fn func(DadJoke)) {
	rcv.ForEachPP(10, fn)
}

func (rcv DadJokeList) ForEachPP(parallelism int, fn func(DadJoke)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv DadJokeList) ForEachPPP(parallelism int, fn func(DadJoke), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan DadJoke, nrJobs)
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
	rcv.ForEach(func(x DadJoke) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv DadJokeList) Count() int {
	return len(rcv)
}

func (rcv DadJokeList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv DadJokeList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv DadJokeList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of DadJoke
func (rcv DadJokeList) ToSlice() []DadJoke {
	return rcv
}

// returns the element at the index
func (rcv DadJokeList) Apply(x int) DadJoke {
	return rcv[x]
}

func (rcv DadJokeList) ApplyOrElse(x int, fn func() DadJoke) DadJoke {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv DadJokeList) FoldLeft(zero DadJoke, fn func(acc DadJoke, x DadJoke) DadJoke) DadJoke {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv DadJokeList) FoldRight(zero DadJoke, fn func(acc DadJoke, x DadJoke) DadJoke) DadJoke {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv DadJokeList) Intersperse(a DadJoke) DadJokeList {
	ys := EmptyDadJokeList()
	rcv.ForEachWithLastFlag(func(last bool, x DadJoke) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv DadJokeList) Fill(num int, a DadJoke) DadJokeList {
	xs := EmptyDadJokeList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv DadJokeList) Tablulate(num int, fn func(int) DadJoke) DadJokeList {
	xs := EmptyDadJokeList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv DadJokeList) Partition(fn func(DadJoke) bool) (DadJokeList, DadJokeList) {
	xs := EmptyDadJokeList()
	ys := EmptyDadJokeList()
	rcv.ForEach(func(x DadJoke) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv DadJokeList) MkString(fn func(DadJoke) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x DadJoke) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv DadJokeList) RangeOf(from int, to int, fn func(int) DadJoke) DadJokeList {
	xs := EmptyDadJokeList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv DadJokeList) Contains(a DadJoke) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv DadJokeList) ContainsNot(a DadJoke) bool {
	return !rcv.Contains(a)
}

func (rcv DadJokeList) Distinct() DadJokeList {
	xs := EmptyDadJokeList()
	rcv.ForEach(func(x DadJoke) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv DadJokeList) Intersect(xs DadJokeList) DadJokeList {
	ys := EmptyDadJokeList()
	rcv.ForEach(func(x DadJoke) {
		xs.ForEach(func(y DadJoke) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv DadJokeList) Slice(from int, to int) DadJokeList {
	return rcv[from : to+1]
}

func (rcv DadJokeList) FlatMapToDadJokeList(fn func(DadJoke) DadJokeList) DadJokeList {
	xs := EmptyDadJokeList()
	rcv.ForEach(func(x DadJoke) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv DadJokeList) MapToDadJoke(fn func(DadJoke) DadJoke) DadJokeList {
	xs := EmptyDadJokeList()
	rcv.ForEach(func(x DadJoke) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv DadJokeList) MapToDadJokeP(mapFn func(DadJoke) DadJoke) DadJokeList {
	return rcv.MapToDadJokePP(10, mapFn)
}

func (rcv DadJokeList) MapToDadJokePP(parallelism int, mapFn func(DadJoke) DadJoke) DadJokeList {
	return rcv.MapToDadJokePPP(parallelism, mapFn, func() {})
}

func (rcv DadJokeList) MapToDadJokePPP(parallelism int, mapFn func(DadJoke) DadJoke, progressFn func()) DadJokeList {
	nrJobs := rcv.Count()
	input := make(chan DadJoke, nrJobs)
	output := make(chan DadJoke, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x DadJoke) {
		input <- x
	})
	close(input)

	xs := EmptyDadJokeList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

// implementation of 'sort.Interface'
func (rcv DadJokeList) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv DadJokeList) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var DadJokeListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv DadJokeList) Less(i, j int) bool {
	return DadJokeListLessFunc(i, j)
}

// i and j are two objects that need to be compared, 
// and based on that comparison the List will be sorted
func (rcv DadJokeList) Sort(fn func(i DadJoke, j DadJoke) bool) DadJokeList {
	DadJokeListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}
