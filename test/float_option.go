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
	FoldToString(zero string, fn func(Float) string) string
	Apply2(a FloatOption, fn func(Float, Float) Float) FloatOption
	Apply3(a FloatOption, b FloatOption, fn func(Float, Float, Float) Float) FloatOption
	Apply4(a FloatOption, b FloatOption, c FloatOption, fn func(Float, Float, Float, Float) Float) FloatOption
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

func (rcv FloatNone) FoldToString(zero string, fn func(Float) string) string {
	return zero
}

func (rcv FloatNone) Apply2(a FloatOption, fn func(Float, Float) Float) FloatOption {
	return noneFloat
}

func (rcv FloatNone) Apply3(a FloatOption, b FloatOption, fn func(Float, Float, Float) Float) FloatOption {
	return noneFloat
}

func (rcv FloatNone) Apply4(a FloatOption, b FloatOption, c FloatOption, fn func(Float, Float, Float, Float) Float) FloatOption {
	return noneFloat
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

func (rcv FloatSome) FoldToString(zero string, fn func(Float) string) string {
	return fn(rcv.Get())
}

func (rcv FloatSome) Apply2(a FloatOption, fn func(Float, Float) Float) FloatOption {
	if rcv.IsDefined() && a.IsDefined() {
		return FloatSome{fn(rcv.Get(), a.Get())}
	} else {
		return noneFloat
	}
}

func (rcv FloatSome) Apply3(a FloatOption, b FloatOption, fn func(Float, Float, Float) Float) FloatOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() {
		return FloatSome{fn(rcv.Get(), a.Get(), b.Get())}
	} else {
		return noneFloat
	}
}

func (rcv FloatSome) Apply4(a FloatOption, b FloatOption, c FloatOption, fn func(Float, Float, Float, Float) Float) FloatOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() && c.IsDefined() {
		return FloatSome{fn(rcv.Get(), a.Get(), b.Get(), c.Get())}
	} else {
		return noneFloat
	}
}
