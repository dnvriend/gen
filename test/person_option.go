// Generated code; DO NOT EDIT.
package test

import (
	"github.com/google/go-cmp/cmp"
)

type PersonOption interface {
	Get() Person
	GetOrElse(fn func() Person) Person
	ForEach(fn func(Person))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a Person) bool
	ContainsNot(a Person) bool
	FoldToString(zero string, fn func(Person) string) string
	Apply2(a PersonOption, fn func(Person, Person) Person) PersonOption
	Apply3(a PersonOption, b PersonOption, fn func(Person, Person, Person) Person) PersonOption
	Apply4(a PersonOption, b PersonOption, c PersonOption, fn func(Person, Person, Person, Person) Person) PersonOption
}

type PersonSome struct {
	a Person
}

type PersonNone struct {
}

var nonePerson = PersonNone{}

func OptionOfPerson(a *Person) PersonOption {
	if a == nil {
		return nonePerson
	} else {
		return PersonSome{*a}
	}
}

// panics when called
func (rcv PersonNone) Get() Person {
	panic("cannot call Get() on None")
}

func (rcv PersonNone) GetOrElse(fn func() Person) Person {
	return fn()
}

func (rcv PersonNone) ForEach(fn func(Person)) {
}

func (rcv PersonNone) IsEmpty() bool {
	return true
}

func (rcv PersonNone) IsNotEmpty() bool {
	return false
}

func (rcv PersonNone) IsDefined() bool {
	return false
}

func (rcv PersonNone) Count() int {
	return 0
}

func (rcv PersonNone) Contains(a Person) bool {
	return false
}

func (rcv PersonNone) ContainsNot(a Person) bool {
	return true
}

func (rcv PersonNone) FoldToString(zero string, fn func(Person) string) string {
	return zero
}

func (rcv PersonNone) Apply2(a PersonOption, fn func(Person, Person) Person) PersonOption {
	return nonePerson
}

func (rcv PersonNone) Apply3(a PersonOption, b PersonOption, fn func(Person, Person, Person) Person) PersonOption {
	return nonePerson
}

func (rcv PersonNone) Apply4(a PersonOption, b PersonOption, c PersonOption, fn func(Person, Person, Person, Person) Person) PersonOption {
	return nonePerson
}

// some
func (rcv PersonSome) Get() Person {
	return rcv.a
}

func (rcv PersonSome) GetOrElse(fn func() Person) Person {
	return rcv.a
}

func (rcv PersonSome) ForEach(fn func(Person)) {
	fn(rcv.a)
}

func (rcv PersonSome) IsEmpty() bool {
	return false
}

func (rcv PersonSome) IsNotEmpty() bool {
	return true
}

// alias
func (rcv PersonSome) IsDefined() bool {
	return true
}

func (rcv PersonSome) Count() int {
	return 1
}

func (rcv PersonSome) Contains(a Person) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv PersonSome) ContainsNot(a Person) bool {
	return !rcv.Contains(a)
}

func (rcv PersonSome) FoldToString(zero string, fn func(Person) string) string {
	return fn(rcv.Get())
}

func (rcv PersonSome) Apply2(a PersonOption, fn func(Person, Person) Person) PersonOption {
	if rcv.IsDefined() && a.IsDefined() {
		return PersonSome{fn(rcv.Get(), a.Get())}
	} else {
		return nonePerson
	}
}

func (rcv PersonSome) Apply3(a PersonOption, b PersonOption, fn func(Person, Person, Person) Person) PersonOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() {
		return PersonSome{fn(rcv.Get(), a.Get(), b.Get())}
	} else {
		return nonePerson
	}
}

func (rcv PersonSome) Apply4(a PersonOption, b PersonOption, c PersonOption, fn func(Person, Person, Person, Person) Person) PersonOption {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() && c.IsDefined() {
		return PersonSome{fn(rcv.Get(), a.Get(), b.Get(), c.Get())}
	} else {
		return nonePerson
	}
}
