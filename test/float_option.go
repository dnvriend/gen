// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
	
)

type FloatOption interface {
	Get() Float
	GetOrElse(fn func() Float) Float
	ForEach(fn func(Float))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a Float) bool
	ContainsNot(a Float) bool
}

type FloatSome struct {
	a Float
}

type FloatNone struct {
}

var noneFloat = FloatNone{}

func OptionOfFloat(a *Float) FloatOption {
	if a == nil {
		return noneFloat
	} else {
		return FloatSome{*a}
	}
}

// panics when called
func (rcv FloatNone) Get() Float {
	panic("cannot call Get() on None")
}

func (rcv FloatNone) GetOrElse(fn func() Float) Float {
	return fn()
}

func (rcv FloatNone) ForEach(fn func(Float)) {	
}

func (rcv FloatNone) IsEmpty() bool {
	return true
}

func (rcv FloatNone) IsNotEmpty() bool {
	return false
}

func (rcv FloatNone) IsDefined() bool {
	return false
}

func (rcv FloatNone) Count() int {
	return 0
}

func (rcv FloatNone) Contains(a Float) bool {
	return false
}

func (rcv FloatNone) ContainsNot(a Float) bool {
	return true
}

// some
func (rcv FloatSome) Get() Float {
	return rcv.a
}

func (rcv FloatSome) GetOrElse(fn func() Float) Float {
	return rcv.a
}

func (rcv FloatSome) ForEach(fn func(Float)) {
	fn(rcv.a)
}

func (rcv FloatSome) IsEmpty() bool {
	return false
}

func (rcv FloatSome) IsNotEmpty() bool {
	return true
}

// alias 
func (rcv FloatSome) IsDefined() bool {
	return true
}

func (rcv FloatSome) Count() int {
	return 1
}

func (rcv FloatSome) Contains(a Float) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv FloatSome) ContainsNot(a Float) bool {
	return !rcv.Contains(a)
}
