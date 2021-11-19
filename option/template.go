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
`))

var importsTemplate = template.Must(template.New("generated").Parse(`
import (
	{{range $index, $import := .Imports}}"{{$import}}"
	{{end}}
)
`))
