// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"strings"
)

type StringList []string

func EmptyStringList() StringList {
	return make([]string, 0)
}

// append an element to the end of the list
func (rcv StringList) Append(x string) StringList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv StringList) AppendAll(xs ...string) StringList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv StringList) Concat(xs ...string) StringList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv StringList) AppendSlice(xs []string) StringList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv StringList) ConcatSlice(xs []string) StringList {
	return rcv.AppendSlice(xs)
}

func (rcv StringList) Reverse() StringList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv StringList) Head() string {
	return rcv[0]
}

func (rcv StringList) HeadOption() StringOption {
	if len(rcv) == 0 {
		return OptionOfString(nil)
	}
	return OptionOfString(&rcv[0])
}

func (rcv StringList) Last() string {
	return rcv[len(rcv)-1]
}

// returns the initial part of the collection, without the last element
func (rcv StringList) Init() StringList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv StringList) Tail() StringList {
	return rcv[1:]
}

// Selects all elements of this list which satisfy a predicate.
func (rcv StringList) Filter(fn func(string) bool) StringList {
	ys := EmptyStringList()
	rcv.ForEach(func(v string) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv StringList) TakeWhile(fn func(string) bool) StringList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv StringList) FilterNot(fn func(string) bool) StringList {
	return rcv.Filter(func(x string) bool { return !fn(x) })
}

// alias for FilterNot
func (rcv StringList) DropWhile(fn func(string) bool) StringList {
	return rcv.FilterNot(fn)
}

func (rcv StringList) ForEach(fn func(string)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv StringList) ForEachWithIndex(fn func(int, string)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv StringList) ForEachWithLastFlag(fn func(bool, string)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv StringList) ForEachP(fn func(string)) {
	rcv.ForEachPP(10, fn)
}

func (rcv StringList) ForEachPP(parallelism int, fn func(string)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv StringList) ForEachPPP(parallelism int, fn func(string), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan string, nrJobs)
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
	rcv.ForEach(func(x string) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv StringList) Count() int {
	return len(rcv)
}

func (rcv StringList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv StringList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv StringList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of string
func (rcv StringList) ToSlice() []string {
	return rcv
}

// returns the element at the index
func (rcv StringList) Apply(x int) string {
	return rcv[x]
}

func (rcv StringList) ApplyOrElse(x int, fn func() string) string {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv StringList) FoldLeft(zero string, fn func(acc string, x string) string) string {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv StringList) FoldRight(zero string, fn func(acc string, x string) string) string {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv StringList) Intersperse(a string) StringList {
	ys := EmptyStringList()
	rcv.ForEachWithLastFlag(func(last bool, x string) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv StringList) Fill(num int, a string) StringList {
	xs := EmptyStringList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv StringList) Tablulate(num int, fn func(int) string) StringList {
	xs := EmptyStringList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv StringList) Partition(fn func(string) bool) (StringList, StringList) {
	xs := EmptyStringList()
	ys := EmptyStringList()
	rcv.ForEach(func(x string) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv StringList) MkString(fn func(string) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x string) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv StringList) RangeOf(from int, to int, fn func(int) string) StringList {
	xs := EmptyStringList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv StringList) Contains(a string) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv StringList) ContainsNot(a string) bool {
	return !rcv.Contains(a)
}

func (rcv StringList) Distinct() StringList {
	xs := EmptyStringList()
	rcv.ForEach(func(x string) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv StringList) Intersect(xs StringList) StringList {
	ys := EmptyStringList()
	rcv.ForEach(func(x string) {
		xs.ForEach(func(y string) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv StringList) Slice(from int, to int) StringList {
	return rcv[from : to+1]
}

func (rcv StringList) FlatMapToStringList(fn func(string) StringList) StringList {
	xs := EmptyStringList()
	rcv.ForEach(func(x string) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv StringList) MapToString(fn func(string) string) StringList {
	xs := EmptyStringList()
	rcv.ForEach(func(x string) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv StringList) MapToStringP(mapFn func(string) string) StringList {
	return rcv.MapToStringPP(10, mapFn)
}

func (rcv StringList) MapToStringPP(parallelism int, mapFn func(string) string) StringList {
	return rcv.MapToStringPPP(parallelism, mapFn, func() {})
}

func (rcv StringList) MapToStringPPP(parallelism int, mapFn func(string) string, progressFn func()) StringList {
	nrJobs := rcv.Count()
	input := make(chan string, nrJobs)
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
	rcv.ForEach(func(x string) {
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

// implementation of 'sort.Interface'
func (rcv StringList) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv StringList) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var StringListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv StringList) Less(i, j int) bool {
	return StringListLessFunc(i, j)
}

// i and j are two objects that need to be compared,
// and based on that comparison the List will be sorted
func (rcv StringList) Sort(fn func(i string, j string) bool) StringList {
	StringListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}

func (rcv StringList) MapToInt(fn func(string) int) IntList {
	ys := make([]int, 0)
	for _, x := range rcv {
		ys = append(ys, fn(x))
	}
	return ys
}

func (rcv StringList) MapToIntWithIndex(fn func(int, string) int) IntList {
	ys := make([]int, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i, x))
	}
	return ys
}

func (rcv StringList) MapToIntWithLastFlag(fn func(bool, string) int) IntList {
	ys := make([]int, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i+1 == len(rcv), x))
	}
	return ys
}

func (rcv StringList) MapToIntP(mapFn func(string) int) IntList {
	return rcv.MapToIntPP(10, mapFn)
}

func (rcv StringList) MapToIntPP(parallelism int, mapFn func(string) int) IntList {
	return rcv.MapToIntPPP(parallelism, mapFn, func() {})
}

func (rcv StringList) MapToIntPPP(parallelism int, mapFn func(string) int, progressFn func()) IntList {
	nrJobs := rcv.Count()
	input := make(chan string, nrJobs)
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
	rcv.ForEach(func(x string) {
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

func (rcv StringList) FlatMapToCatList(fn func(string) CatList) CatList {
	xs := EmptyCatList()
	rcv.ForEach(func(x string) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

// joins using the character and returns the string
func (rcv StringList) Join(sep string) String {
	return rcv.Intersperse(sep).MkString(func(x string) string { return x })
}
