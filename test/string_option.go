// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
)

type StringOption interface {
	Get() string
	GetOrElse(fn func() string) string
	ForEach(fn func(string))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a string) bool
	ContainsNot(a string) bool
	FoldToString(zero string, fn func(string) string) string
	Apply2(a StringOption, fn func(string, string) string) StringOption
	Apply3(a StringOption, b StringOption, fn func(string, string, string) string) StringOption
	Apply4(a StringOption, b StringOption, c StringOption, fn func(string, string, string, string) string) StringOption
}

type StringSome struct {
	a string
}

type StringNone struct {
}

var noneString = StringNone{}

func OptionOfString(a *string) StringOption {
	if a == nil {
		return noneString
	} else {
		return StringSome{*a}
	}
}

// panics when called
func (rcv StringNone) Get() string {
	panic("cannot call Get() on None")
}

func (rcv StringNone) GetOrElse(fn func() string) string {
	return fn()
}

func (rcv StringNone) ForEach(fn func(string)) {
}

func (rcv StringNone) IsEmpty() bool {
	return true
}

func (rcv StringNone) IsNotEmpty() bool {
	return false
}

func (rcv StringNone) IsDefined() bool {
	return false
}

func (rcv StringNone) Count() int {
	return 0
}

func (rcv StringNone) Contains(a string) bool {
	return false
}

func (rcv StringNone) ContainsNot(a string) bool {
	return true
}

func (rcv StringNone) FoldToString(zero string, fn func(string) string) string {
	return zero
}

func (rcv StringNone) Apply2(a StringOption, fn func(string, string) string) StringOption {
	return noneString
}

func (rcv StringNone) Apply3(a StringOption, b StringOption, fn func(string, string, string) string) StringOption {
	return noneString
}

func (rcv StringNone) Apply4(a StringOption, b StringOption, c StringOption, fn func(string, string, string, string) string) StringOption {
	return noneString
}

// some
func (rcv StringSome) Get() string {
	return rcv.a
}

func (rcv StringSome) GetOrElse(fn func() string) string {
	return rcv.a
}

func (rcv StringSome) ForEach(fn func(string)) {
	fn(rcv.a)
}

func (rcv StringSome) IsEmpty() bool {
	return false
}

func (rcv StringSome) IsNotEmpty() bool {
	return true
}

// alias
func (rcv StringSome) IsDefined() bool {
	return true
}

func (rcv StringSome) Count() int {
	return 1
}

func (rcv StringSome) Contains(a string) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv StringSome) ContainsNot(a string) bool {
	return !rcv.Contains(a)
}

func (rcv StringSome) FoldToString(zero string, fn func(string) string) string {
	return fn(rcv.Get())
}

func (rcv StringSome) Apply2(a StringOption, fn func(string, string) string) StringOption {
	if rcv.IsDefined() && a.IsDefined() {
		return StringSome{fn(rcv.Get(), a.Get())}
	} else {
		return noneString
	}
}

func (rcv StringSome) Apply3(a StringOption, b StringOption, fn func(string, string, string) string) StringOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() {
		return StringSome{fn(rcv.Get(), a.Get(), b.Get())}
	} else {
		return noneString
	}
}

func (rcv StringSome) Apply4(a StringOption, b StringOption, c StringOption, fn func(string, string, string, string) string) StringOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() && c.IsDefined() {
		return StringSome{fn(rcv.Get(), a.Get(), b.Get(), c.Get())}
	} else {
		return noneString
	}
}
