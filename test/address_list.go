// Generated code; DO NOT EDIT.
package test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
)

type AddressList []Address

func EmptyAddressList() AddressList {
	return make([]Address, 0)
}

// append an element to the end of the list
func (rcv AddressList) Append(x Address) AddressList {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv AddressList) AppendAll(xs ...Address) AddressList {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv AddressList) Concat(xs ...Address) AddressList {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv AddressList) AppendSlice(xs []Address) AddressList {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv AddressList) ConcatSlice(xs []Address) AddressList {
	return rcv.AppendSlice(xs)
}

func (rcv AddressList) Reverse() AddressList {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv AddressList) Head() Address {
	return rcv[0]
}

func (rcv AddressList) HeadOption() AddressOption {
	if len(rcv) == 0 {
		return OptionOfAddress(nil)
	}
	return OptionOfAddress(&rcv[0])
}

func (rcv AddressList) Last() Address {
	return rcv[len(rcv)-1]
}

// returns the initial part of the collection, without the last element
func (rcv AddressList) Init() AddressList {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv AddressList) Tail() AddressList {
	return rcv[1:]
}

// Selects all elements of this list which satisfy a predicate.
func (rcv AddressList) Filter(fn func(Address) bool) AddressList {
	ys := EmptyAddressList()
	rcv.ForEach(func(v Address) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv AddressList) TakeWhile(fn func(Address) bool) AddressList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv AddressList) FilterNot(fn func(Address) bool) AddressList {
	return rcv.Filter(func(x Address) bool { return !fn(x) })
}

// alias for FilterNot
func (rcv AddressList) DropWhile(fn func(Address) bool) AddressList {
	return rcv.FilterNot(fn)
}

func (rcv AddressList) ForEach(fn func(Address)) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv AddressList) ForEachWithIndex(fn func(int, Address)) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv AddressList) ForEachWithLastFlag(fn func(bool, Address)) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv AddressList) ForEachP(fn func(Address)) {
	rcv.ForEachPP(10, fn)
}

func (rcv AddressList) ForEachPP(parallelism int, fn func(Address)) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv AddressList) ForEachPPP(parallelism int, fn func(Address), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan Address, nrJobs)
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
	rcv.ForEach(func(x Address) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv AddressList) Count() int {
	return len(rcv)
}

func (rcv AddressList) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv AddressList) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv AddressList) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of Address
func (rcv AddressList) ToSlice() []Address {
	return rcv
}

// returns the element at the index
func (rcv AddressList) Apply(x int) Address {
	return rcv[x]
}

func (rcv AddressList) ApplyOrElse(x int, fn func() Address) Address {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv AddressList) FoldLeft(zero Address, fn func(acc Address, x Address) Address) Address {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv AddressList) FoldRight(zero Address, fn func(acc Address, x Address) Address) Address {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv AddressList) Intersperse(a Address) AddressList {
	ys := EmptyAddressList()
	rcv.ForEachWithLastFlag(func(last bool, x Address) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv AddressList) Fill(num int, a Address) AddressList {
	xs := EmptyAddressList()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv AddressList) Tablulate(num int, fn func(int) Address) AddressList {
	xs := EmptyAddressList()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv AddressList) Partition(fn func(Address) bool) (AddressList, AddressList) {
	xs := EmptyAddressList()
	ys := EmptyAddressList()
	rcv.ForEach(func(x Address) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv AddressList) MkString() String {
	var builder strings.Builder
	rcv.ForEach(func(x Address) {
		builder.WriteString(fmt.Sprintf("%v", x))
	})
	return String(builder.String())
}

func (rcv AddressList) RangeOf(from int, to int, fn func(int) Address) AddressList {
	xs := EmptyAddressList()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv AddressList) Contains(a Address) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv AddressList) ContainsNot(a Address) bool {
	return !rcv.Contains(a)
}

func (rcv AddressList) Distinct() AddressList {
	xs := EmptyAddressList()
	rcv.ForEach(func(x Address) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv AddressList) Intersect(xs AddressList) AddressList {
	ys := EmptyAddressList()
	rcv.ForEach(func(x Address) {
		xs.ForEach(func(y Address) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv AddressList) Slice(from int, to int) AddressList {
	return rcv[from : to+1]
}

func (rcv AddressList) FlatMapToAddressList(fn func(Address) AddressList) AddressList {
	xs := EmptyAddressList()
	rcv.ForEach(func(x Address) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv AddressList) MapToAddress(fn func(Address) Address) AddressList {
	xs := EmptyAddressList()
	rcv.ForEach(func(x Address) {
		xs = xs.Append(fn(x))
	})
	return xs
}
