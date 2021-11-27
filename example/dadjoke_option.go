// Generated code; DO NOT EDIT.
package main

import (
	"github.com/google/go-cmp/cmp"
	
)

type DadJokeOption interface {
	Get() DadJoke
	GetOrElse(fn func() DadJoke) DadJoke
	ForEach(fn func(DadJoke))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a DadJoke) bool
	ContainsNot(a DadJoke) bool
}

type DadJokeSome struct {
	a DadJoke
}

type DadJokeNone struct {
}

var noneDadJoke = DadJokeNone{}

func OptionOfDadJoke(a *DadJoke) DadJokeOption {
	if a == nil {
		return noneDadJoke
	} else {
		return DadJokeSome{*a}
	}
}

// panics when called
func (rcv DadJokeNone) Get() DadJoke {
	panic("cannot call Get() on None")
}

func (rcv DadJokeNone) GetOrElse(fn func() DadJoke) DadJoke {
	return fn()
}

func (rcv DadJokeNone) ForEach(fn func(DadJoke)) {	
}

func (rcv DadJokeNone) IsEmpty() bool {
	return true
}

func (rcv DadJokeNone) IsNotEmpty() bool {
	return false
}

func (rcv DadJokeNone) IsDefined() bool {
	return false
}

func (rcv DadJokeNone) Count() int {
	return 0
}

func (rcv DadJokeNone) Contains(a DadJoke) bool {
	return false
}

func (rcv DadJokeNone) ContainsNot(a DadJoke) bool {
	return true
}

// some
func (rcv DadJokeSome) Get() DadJoke {
	return rcv.a
}

func (rcv DadJokeSome) GetOrElse(fn func() DadJoke) DadJoke {
	return rcv.a
}

func (rcv DadJokeSome) ForEach(fn func(DadJoke)) {
	fn(rcv.a)
}

func (rcv DadJokeSome) IsEmpty() bool {
	return false
}

func (rcv DadJokeSome) IsNotEmpty() bool {
	return true
}

// alias 
func (rcv DadJokeSome) IsDefined() bool {
	return true
}

func (rcv DadJokeSome) Count() int {
	return 1
}

func (rcv DadJokeSome) Contains(a DadJoke) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv DadJokeSome) ContainsNot(a DadJoke) bool {
	return !rcv.Contains(a)
}
