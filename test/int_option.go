// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
)

type IntOption interface {
	Get() int
	GetOrElse(fn func() int) int
	ForEach(fn func(int))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a int) bool
	ContainsNot(a int) bool
}

type IntSome struct {
	a int
}

type IntNone struct {
}

var noneInt = IntNone{}

func OptionOfInt(a *int) IntOption {
	if a == nil {
		return noneInt
	} else {
		return IntSome{*a}
	}
}

// panics when called
func (rcv IntNone) Get() int {
	panic("cannot call Get() on None")
}

func (rcv IntNone) GetOrElse(fn func() int) int {
	return fn()
}

func (rcv IntNone) ForEach(fn func(int)) {
}

func (rcv IntNone) IsEmpty() bool {
	return true
}

func (rcv IntNone) IsNotEmpty() bool {
	return false
}

func (rcv IntNone) IsDefined() bool {
	return false
}

func (rcv IntNone) Count() int {
	return 0
}

func (rcv IntNone) Contains(a int) bool {
	return false
}

func (rcv IntNone) ContainsNot(a int) bool {
	return true
}

// some
func (rcv IntSome) Get() int {
	return rcv.a
}

func (rcv IntSome) GetOrElse(fn func() int) int {
	return rcv.a
}

func (rcv IntSome) ForEach(fn func(int)) {
	fn(rcv.a)
}

func (rcv IntSome) IsEmpty() bool {
	return false
}

func (rcv IntSome) IsNotEmpty() bool {
	return true
}

// alias
func (rcv IntSome) IsDefined() bool {
	return true
}

func (rcv IntSome) Count() int {
	return 1
}

func (rcv IntSome) Contains(a int) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv IntSome) ContainsNot(a int) bool {
	return !rcv.Contains(a)
}
