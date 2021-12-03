// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
)

type AddressOption interface {
	Get() Address
	GetOrElse(fn func() Address) Address
	ForEach(fn func(Address))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a Address) bool
	ContainsNot(a Address) bool
	FoldToString(zero string, fn func(Address) string) string
}

type AddressSome struct {
	a Address
}

type AddressNone struct {
}

var noneAddress = AddressNone{}

func OptionOfAddress(a *Address) AddressOption {
	if a == nil {
		return noneAddress
	} else {
		return AddressSome{*a}
	}
}

// panics when called
func (rcv AddressNone) Get() Address {
	panic("cannot call Get() on None")
}

func (rcv AddressNone) GetOrElse(fn func() Address) Address {
	return fn()
}

func (rcv AddressNone) ForEach(fn func(Address)) {
}

func (rcv AddressNone) IsEmpty() bool {
	return true
}

func (rcv AddressNone) IsNotEmpty() bool {
	return false
}

func (rcv AddressNone) IsDefined() bool {
	return false
}

func (rcv AddressNone) Count() int {
	return 0
}

func (rcv AddressNone) Contains(a Address) bool {
	return false
}

func (rcv AddressNone) ContainsNot(a Address) bool {
	return true
}

func (rcv AddressNone) FoldToString(zero string, fn func(Address) string) string {
	return zero
}

// some
func (rcv AddressSome) Get() Address {
	return rcv.a
}

func (rcv AddressSome) GetOrElse(fn func() Address) Address {
	return rcv.a
}

func (rcv AddressSome) ForEach(fn func(Address)) {
	fn(rcv.a)
}

func (rcv AddressSome) IsEmpty() bool {
	return false
}

func (rcv AddressSome) IsNotEmpty() bool {
	return true
}

// alias
func (rcv AddressSome) IsDefined() bool {
	return true
}

func (rcv AddressSome) Count() int {
	return 1
}

func (rcv AddressSome) Contains(a Address) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv AddressSome) ContainsNot(a Address) bool {
	return !rcv.Contains(a)
}

func (rcv AddressSome) FoldToString(zero string, fn func(Address) string) string {
	return fn(rcv.Get())
}
