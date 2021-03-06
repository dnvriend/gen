// Generated code; DO NOT EDIT.
package main

import (
	"sort"
	"strings"
	"github.com/google/go-cmp/cmp"
	
)

type PersonList []Person

func EmptyPersonList() PersonList {
	return make([]Person, 0)
}

// append an element to the end of the list
func (rcv PersonList) Append(x Person) PersonList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv PersonList) AppendAll(xs ...Person) PersonList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv PersonList) Concat(xs ...Person) PersonList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv PersonList) AppendSlice(xs []Person) PersonList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv PersonList) ConcatSlice(xs []Person) PersonList {
	return rcv.AppendSlice(xs)
}

func (rcv PersonList) Reverse() PersonList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv PersonList) Head() Person {
	return rcv[0] 
}

func (rcv PersonList) HeadOption() PersonOption {
	if len(rcv) == 0 {
		return OptionOfPerson(nil)
	} 
	return OptionOfPerson(&rcv[0])
}

func (rcv PersonList) Last() Person {
	return rcv[len(rcv)-1] 
}

// returns the initial part of the collection, without the last element
func (rcv PersonList) Init() PersonList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv PersonList) Tail() PersonList {
	return rcv[1:]
} 

// Selects all elements of this list which satisfy a predicate.
func (rcv PersonList) Filter(fn func(Person) bool) PersonList {
	ys := EmptyPersonList()
 	rcv.ForEach(func(v Person) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv PersonList) TakeWhile(fn func(Person) bool) PersonList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv PersonList) FilterNot(fn func(Person) bool) PersonList {
	return rcv.Filter(func (x Person) bool { return !fn(x)})
}

// alias for FilterNot
func (rcv PersonList) DropWhile(fn func(Person) bool) PersonList {
	return rcv.FilterNot(fn)
}

func (rcv PersonList) ForEach(fn func(Person)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv PersonList) ForEachWithIndex(fn func(int, Person)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv PersonList) ForEachWithLastFlag(fn func(bool, Person)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv PersonList) ForEachP(fn func(Person)) {
	rcv.ForEachPP(10, fn)
}

func (rcv PersonList) ForEachPP(parallelism int, fn func(Person)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv PersonList) ForEachPPP(parallelism int, fn func(Person), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan Person, nrJobs)
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
	rcv.ForEach(func(x Person) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv PersonList) Count() int {
	return len(rcv)
}

func (rcv PersonList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv PersonList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv PersonList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of Person
func (rcv PersonList) ToSlice() []Person {
	return rcv
}

// returns the element at the index
func (rcv PersonList) Apply(x int) Person {
	return rcv[x]
}

func (rcv PersonList) ApplyOrElse(x int, fn func() Person) Person {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv PersonList) FoldLeft(zero Person, fn func(acc Person, x Person) Person) Person {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv PersonList) FoldRight(zero Person, fn func(acc Person, x Person) Person) Person {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv PersonList) Intersperse(a Person) PersonList {
	ys := EmptyPersonList()
	rcv.ForEachWithLastFlag(func(last bool, x Person) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv PersonList) Fill(num int, a Person) PersonList {
	xs := EmptyPersonList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv PersonList) Tablulate(num int, fn func(int) Person) PersonList {
	xs := EmptyPersonList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv PersonList) Partition(fn func(Person) bool) (PersonList, PersonList) {
	xs := EmptyPersonList()
	ys := EmptyPersonList()
	rcv.ForEach(func(x Person) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv PersonList) MkString(fn func(Person) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x Person) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv PersonList) RangeOf(from int, to int, fn func(int) Person) PersonList {
	xs := EmptyPersonList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv PersonList) Contains(a Person) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv PersonList) ContainsNot(a Person) bool {
	return !rcv.Contains(a)
}

func (rcv PersonList) Distinct() PersonList {
	xs := EmptyPersonList()
	rcv.ForEach(func(x Person) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv PersonList) Intersect(xs PersonList) PersonList {
	ys := EmptyPersonList()
	rcv.ForEach(func(x Person) {
		xs.ForEach(func(y Person) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv PersonList) Slice(from int, to int) PersonList {
	return rcv[from : to+1]
}

func (rcv PersonList) FlatMapToPersonList(fn func(Person) PersonList) PersonList {
	xs := EmptyPersonList()
	rcv.ForEach(func(x Person) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv PersonList) MapToPerson(fn func(Person) Person) PersonList {
	xs := EmptyPersonList()
	rcv.ForEach(func(x Person) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv PersonList) MapToPersonP(mapFn func(Person) Person) PersonList {
	return rcv.MapToPersonPP(10, mapFn)
}

func (rcv PersonList) MapToPersonPP(parallelism int, mapFn func(Person) Person) PersonList {
	return rcv.MapToPersonPPP(parallelism, mapFn, func() {})
}

func (rcv PersonList) MapToPersonPPP(parallelism int, mapFn func(Person) Person, progressFn func()) PersonList {
	nrJobs := rcv.Count()
	input := make(chan Person, nrJobs)
	output := make(chan Person, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x Person) {
		input <- x
	})
	close(input)

	xs := EmptyPersonList()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

// implementation of 'sort.Interface'
func (rcv PersonList) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv PersonList) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var PersonListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv PersonList) Less(i, j int) bool {
	return PersonListLessFunc(i, j)
}

// i and j are two objects that need to be compared, 
// and based on that comparison the List will be sorted
func (rcv PersonList) Sort(fn func(i Person, j Person) bool) PersonList {
	PersonListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}
