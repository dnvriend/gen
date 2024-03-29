package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:generate gen option -t int -p test
//go:generate gen option -t string -p test
//go:generate gen option -t Person -p test
//go:generate gen option -t Address -p test
//go:generate gen option -t Cat -p test

func TestGet(t *testing.T) {
	a := 1
	x := OptionOfInt(&a)
	assert.Equal(t, 1, x.Get())
	assert.Equal(t, 1, x.GetOrElse(func() int { return 2 }))

	y := OptionOfString(nil)
	assert.Equal(t, noneString, y)
	assert.Equal(t, noneString.GetOrElse(func() string { return "bla" }), "bla")

	str := "foo"
	z := OptionOfString(&str)
	assert.Equal(t, StringSome{"foo"}, z)
	assert.Equal(t, z.Get(), "foo")
	assert.Equal(t, z.GetOrElse(func() string { return "bla" }), "foo")

	p := Person{
		Name:      "mr black",
		Age:       47,
		Addresses: EmptyAddressList().Append(Address{}),
		Cats:      EmptyCatList().Append(Cat{}),
	}
	po := OptionOfPerson(&p)
	assert.Equal(t, PersonSome{p}, po)
	assert.Equal(t, po.Get(), p)
	assert.Equal(t, po.GetOrElse(func() Person { return Person{} }), p)
}

func TestSome(t *testing.T) {
	cat := Cat{}
	o := OptionOfCat(&cat)
	assert.Equal(t, CatSome{cat}, o)
	assert.Equal(t, 1, o.Count())
	assert.Equal(t, false, o.IsEmpty())
	assert.Equal(t, true, o.IsNotEmpty())
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, true, o.Contains(cat))
	assert.Equal(t, false, o.ContainsNot(cat))
}

func TestNone(t *testing.T) {
	cat := Cat{}
	o := OptionOfCat(nil)
	assert.Equal(t, noneCat, o)
	assert.Equal(t, 0, o.Count())
	assert.Equal(t, true, o.IsEmpty())
	assert.Equal(t, false, o.IsNotEmpty())
	assert.Equal(t, false, o.IsDefined())
	assert.Equal(t, false, o.Contains(cat))
	assert.Equal(t, true, o.ContainsNot(cat))
}

func TestApply234(t *testing.T) {
	foo := "foo"
	bar := "bar"
	baz := "baz"
	quz := "quz"
	o_foo := OptionOfString(&foo)
	o_bar := OptionOfString(&bar)
	o_baz := OptionOfString(&baz)
	o_quz := OptionOfString(&quz)

	res := o_foo.Apply2(o_bar, func(s string, s2 string) string {
		return fmt.Sprintf("%s-%s", s, s2)
	})
	foobar := "foo-bar"
	assert.Equal(t, OptionOfString(&foobar), res)

	res = o_foo.Apply3(o_bar, o_baz, func(s string, s2 string, s3 string) string {
		return fmt.Sprintf("%s-%s-%s", s, s2, s3)
	})
	foobarbaz := "foo-bar-baz"
	assert.Equal(t, OptionOfString(&foobarbaz), res)

	res = o_foo.Apply4(o_bar, o_baz, o_quz, func(s string, s2 string, s3 string, s4 string) string {
		return fmt.Sprintf("%s-%s-%s-%s", s, s2, s3, s4)
	})
	foobarbazquz := "foo-bar-baz-quz"
	assert.Equal(t, OptionOfString(&foobarbazquz), res)
}
