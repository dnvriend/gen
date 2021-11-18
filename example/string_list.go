// Generated code; DO NOT EDIT.
package main

import (
	"errors"
	"fmt"
	"strings"
	"github.com/google/go-cmp/cmp"
	
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

func (rcv StringList) Head() string {
	return rcv[0] 
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
	ys := make([]string, 0)
	for _, v := range rcv {
		if fn(v) {
			ys = append(ys, v)
		}
	}
	return ys
}

// alias for Filter
func (rcv StringList) TakeWhile(fn func(string) bool) StringList {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv StringList) FilterNot(fn func(string) bool) StringList {
	ys := make([]string, 0)
	for _, v := range rcv {
		if !fn(v) {
			ys = append(ys, v)
		}
	}
	return ys
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

// Finds the first element of the list satisfying a predicate, if any.
func (rcv StringList) Find(fn func(string) bool) (*string, error) {
	for _, x := range rcv {
		if fn(x) {
			return &x, nil
		}
	}
	return nil, errors.New("Could not find element")
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

func (rcv StringList) MkString() string {
	var builder strings.Builder
	rcv.ForEach(func(x string) {
		builder.WriteString(fmt.Sprintf("%v", x))
	})
	return builder.String()
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
