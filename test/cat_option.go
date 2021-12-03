// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
)

type CatOption interface {
	Get() Cat
	GetOrElse(fn func() Cat) Cat
	ForEach(fn func(Cat))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a Cat) bool
	ContainsNot(a Cat) bool
	FoldToString(zero string, fn func(Cat) string) string
}

type CatSome struct {
	a Cat
}

type CatNone struct {
}

var noneCat = CatNone{}

func OptionOfCat(a *Cat) CatOption {
	if a == nil {
		return noneCat
	} else {
		return CatSome{*a}
	}
}

// panics when called
func (rcv CatNone) Get() Cat {
	panic("cannot call Get() on None")
}

func (rcv CatNone) GetOrElse(fn func() Cat) Cat {
	return fn()
}

func (rcv CatNone) ForEach(fn func(Cat)) {
}

func (rcv CatNone) IsEmpty() bool {
	return true
}

func (rcv CatNone) IsNotEmpty() bool {
	return false
}

func (rcv CatNone) IsDefined() bool {
	return false
}

func (rcv CatNone) Count() int {
	return 0
}

func (rcv CatNone) Contains(a Cat) bool {
	return false
}

func (rcv CatNone) ContainsNot(a Cat) bool {
	return true
}

func (rcv CatNone) FoldToString(zero string, fn func(Cat) string) string {
	return zero
}

// some
func (rcv CatSome) Get() Cat {
	return rcv.a
}

func (rcv CatSome) GetOrElse(fn func() Cat) Cat {
	return rcv.a
}

func (rcv CatSome) ForEach(fn func(Cat)) {
	fn(rcv.a)
}

func (rcv CatSome) IsEmpty() bool {
	return false
}

func (rcv CatSome) IsNotEmpty() bool {
	return true
}

// alias
func (rcv CatSome) IsDefined() bool {
	return true
}

func (rcv CatSome) Count() int {
	return 1
}

func (rcv CatSome) Contains(a Cat) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv CatSome) ContainsNot(a Cat) bool {
	return !rcv.Contains(a)
}

func (rcv CatSome) FoldToString(zero string, fn func(Cat) string) string {
	return fn(rcv.Get())
}
