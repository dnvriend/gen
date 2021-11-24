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
