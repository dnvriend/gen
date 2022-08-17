package option

import "text/template"

var baseTmpl = template.Must(template.New("generated").Parse(`
type {{.TypeName}}Option interface {
	Get() {{.Type}}
	GetOrElse(fn func() {{.Type}}) {{.Type}}
	ForEach(fn func({{.Type}}))
	IsEmpty() bool
	IsNotEmpty() bool
	IsDefined() bool
	Count() int
	Contains(a {{.Type}}) bool
	ContainsNot(a {{.Type}}) bool
	FoldToString(zero string, fn func({{.Type}}) string) string
	Apply2(a {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option
	Apply3(a {{.TypeName}}Option, b {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option
	Apply4(a {{.TypeName}}Option, b {{.TypeName}}Option, c {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option
}

type {{.TypeName}}Some struct {
	a {{.Type}}
}

type {{.TypeName}}None struct {
}

var none{{.TypeName}} = {{.TypeName}}None{}

func OptionOf{{.TypeName}}(a *{{.Type}}) {{.TypeName}}Option {
	if a == nil {
		return none{{.TypeName}}
	} else {
		return {{.TypeName}}Some{*a}
	}
}

// panics when called
func (rcv {{.TypeName}}None) Get() {{.Type}} {
	panic("cannot call Get() on None")
}

func (rcv {{.TypeName}}None) GetOrElse(fn func() {{.Type}}) {{.Type}} {
	return fn()
}

func (rcv {{.TypeName}}None) ForEach(fn func({{.Type}})) {	
}

func (rcv {{.TypeName}}None) IsEmpty() bool {
	return true
}

func (rcv {{.TypeName}}None) IsNotEmpty() bool {
	return false
}

func (rcv {{.TypeName}}None) IsDefined() bool {
	return false
}

func (rcv {{.TypeName}}None) Count() int {
	return 0
}

func (rcv {{.TypeName}}None) Contains(a {{.Type}}) bool {
	return false
}

func (rcv {{.TypeName}}None) ContainsNot(a {{.Type}}) bool {
	return true
}

func (rcv {{.TypeName}}None) FoldToString(zero string, fn func({{.Type}}) string) string {
	return zero
}

func (rcv {{.TypeName}}None) Apply2(a {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	return none{{.TypeName}}
}

func (rcv {{.TypeName}}None) Apply3(a {{.TypeName}}Option, b {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	return none{{.TypeName}}
}

func (rcv {{.TypeName}}None) Apply4(a {{.TypeName}}Option, b {{.TypeName}}Option, c {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	return none{{.TypeName}}
}

// some
func (rcv {{.TypeName}}Some) Get() {{.Type}} {
	return rcv.a
}

func (rcv {{.TypeName}}Some) GetOrElse(fn func() {{.Type}}) {{.Type}} {
	return rcv.a
}

func (rcv {{.TypeName}}Some) ForEach(fn func({{.Type}})) {
	fn(rcv.a)
}

func (rcv {{.TypeName}}Some) IsEmpty() bool {
	return false
}

func (rcv {{.TypeName}}Some) IsNotEmpty() bool {
	return true
}

// alias 
func (rcv {{.TypeName}}Some) IsDefined() bool {
	return true
}

func (rcv {{.TypeName}}Some) Count() int {
	return 1
}

func (rcv {{.TypeName}}Some) Contains(a {{.Type}}) bool {
	return cmp.Equal(rcv.a, a)
}

func (rcv {{.TypeName}}Some) ContainsNot(a {{.Type}}) bool {
	return !rcv.Contains(a)
}

func (rcv {{.TypeName}}Some) FoldToString(zero string, fn func({{.Type}}) string) string {	
	return fn(rcv.Get()) 	
}

func (rcv {{.TypeName}}Some) Apply2(a {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	if rcv.IsDefined() && a.IsDefined() {
		return {{.TypeName}}Some { fn(rcv.Get(), a.Get()) }
	} else {
		return none{{.TypeName}}
	}	
}

func (rcv {{.TypeName}}Some) Apply3(a {{.TypeName}}Option, b {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() {
		return {{.TypeName}}Some { fn(rcv.Get(), a.Get(), b.Get()) }
	} else {
		return none{{.TypeName}}
	}
}

func (rcv {{.TypeName}}Some) Apply4(a {{.TypeName}}Option, b {{.TypeName}}Option, c {{.TypeName}}Option, fn func({{.Type}}, {{.Type}}, {{.Type}}, {{.Type}}) {{.Type}}) {{.TypeName}}Option {
	if rcv.IsDefined() && a.IsDefined() && b.IsDefined() && c.IsDefined() {
		return {{.TypeName}}Some { fn(rcv.Get(), a.Get(), b.Get(), c.Get()) }
	} else {
		return none{{.TypeName}}
	}
}
`))

var importsTemplate = template.Must(template.New("generated").Parse(`
import (
	{{range $index, $import := .Imports}}"{{$import}}"
	{{end}}
)
`))
