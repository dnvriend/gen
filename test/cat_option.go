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
	Apply2(a CatOption, fn func(Cat, Cat) Cat) CatOption
	Apply3(a CatOption, b CatOption, fn func(Cat, Cat, Cat) Cat) CatOption
	Apply4(a CatOption, b CatOption, c CatOption, fn func(Cat, Cat, Cat, Cat) Cat) CatOption
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

func (rcv CatNone) Apply2(a CatOption, fn func(Cat, Cat) Cat) CatOption {
	return noneCat
}

func (rcv CatNone) Apply3(a CatOption, b CatOption, fn func(Cat, Cat, Cat) Cat) CatOption {
	return noneCat
}

func (rcv CatNone) Apply4(a CatOption, b CatOption, c CatOption, fn func(Cat, Cat, Cat, Cat) Cat) CatOption {
	return noneCat
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

func (rcv CatSome) Apply2(a CatOption, fn func(Cat, Cat) Cat) CatOption {
	if rcv.IsDefined() && a.IsDefined() {
		return CatSome{fn(rcv.Get(), a.Get())}
	} else {
		return noneCat
	}
}

func (rcv CatSome) Apply3(a CatOption, b CatOption, fn func(Cat, Cat, Cat) Cat) CatOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() {
		return CatSome{fn(rcv.Get(), a.Get(), b.Get())}
	} else {
		return noneCat
	}
}

func (rcv CatSome) Apply4(a CatOption, b CatOption, c CatOption, fn func(Cat, Cat, Cat, Cat) Cat) CatOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() && c.IsDefined() {
		return CatSome{fn(rcv.Get(), a.Get(), b.Get(), c.Get())}
	} else {
		return noneCat
	}
}
