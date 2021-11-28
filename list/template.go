package list

import (
	"text/template"
)

var baseTmpl = template.Must(template.New("generated").Parse(`
type {{.TypeName}}List []{{.Type}}

func Empty{{.TypeName}}List() {{.TypeName}}List {
	return make([]{{.Type}}, 0)
}

// append an element to the end of the list
func (rcv {{.TypeName}}List) Append(x {{.Type}}) {{.TypeName}}List {
	return append(rcv, x)
}

// append elements to the end of the list
func (rcv {{.TypeName}}List) AppendAll(xs ...{{.Type}}) {{.TypeName}}List {
	return append(rcv, xs...)
}

// alias for AppendAll
func (rcv {{.TypeName}}List) Concat(xs ...{{.Type}}) {{.TypeName}}List {
	return rcv.AppendAll(xs...)
}

// append a slice to the end of the list
func (rcv {{.TypeName}}List) AppendSlice(xs []{{.Type}}) {{.TypeName}}List {
	return append(rcv, xs...)
}

// alias for AppendSlice
func (rcv {{.TypeName}}List) ConcatSlice(xs []{{.Type}}) {{.TypeName}}List {
	return rcv.AppendSlice(xs)
}

func (rcv {{.TypeName}}List) Reverse() {{.TypeName}}List {
	for i, j := 0, len(rcv)-1; i < j; i, j = i+1, j-1 {
		rcv[i], rcv[j] = rcv[j], rcv[i]
	}
	return rcv
}

// panics when the list is empty
func (rcv {{.TypeName}}List) Head() {{.Type}} {
	return rcv[0] 
}

func (rcv {{.TypeName}}List) HeadOption() {{.TypeName}}Option {
	if len(rcv) == 0 {
		return OptionOf{{.TypeName}}(nil)
	} 
	return OptionOf{{.TypeName}}(&rcv[0])
}

func (rcv {{.TypeName}}List) Last() {{.Type}} {
	return rcv[len(rcv)-1] 
}

// returns the initial part of the collection, without the last element
func (rcv {{.TypeName}}List) Init() {{.TypeName}}List {
	return rcv[:len(rcv)-1]
}

// The rest of the collection without its first element.
func (rcv {{.TypeName}}List) Tail() {{.TypeName}}List {
	return rcv[1:]
} 

// Selects all elements of this list which satisfy a predicate.
func (rcv {{.TypeName}}List) Filter(fn func({{.Type}}) bool) {{.TypeName}}List {
	ys := Empty{{.TypeName}}List()
 	rcv.ForEach(func(v {{.Type}}) {
		if fn(v) {
			ys = ys.Append(v)
		}
	})
	return ys
}

// alias for Filter
func (rcv {{.TypeName}}List) TakeWhile(fn func({{.Type}}) bool) {{.TypeName}}List {
	return rcv.Filter(fn)
}

// Selects all elements of this list which do not satisfy a predicate.
func (rcv {{.TypeName}}List) FilterNot(fn func({{.Type}}) bool) {{.TypeName}}List {
	return rcv.Filter(func (x {{.Type}}) bool { return !fn(x)})
}

// alias for FilterNot
func (rcv {{.TypeName}}List) DropWhile(fn func({{.Type}}) bool) {{.TypeName}}List {
	return rcv.FilterNot(fn)
}

func (rcv {{.TypeName}}List) ForEach(fn func({{.Type}})) {
	for _, x := range rcv {
		fn(x)
	}
}

func (rcv {{.TypeName}}List) ForEachWithIndex(fn func(int, {{.Type}})) {
	for i, x := range rcv {
		fn(i, x)
	}
}

func (rcv {{.TypeName}}List) ForEachWithLastFlag(fn func(bool, {{.Type}})) {
	for i, x := range rcv {
		fn(i+1 == len(rcv), x)
	}
}

func (rcv {{.TypeName}}List) ForEachP(fn func({{.Type}})) {
	rcv.ForEachPP(10, fn)
}

func (rcv {{.TypeName}}List) ForEachPP(parallelism int, fn func({{.Type}})) {
	rcv.ForEachPPP(parallelism, fn, func() {})
}

func (rcv {{.TypeName}}List) ForEachPPP(parallelism int, fn func({{.Type}}), progressFn func()) {
	nrJobs := rcv.Count()
	input := make(chan {{.Type}}, nrJobs)
	output := make(chan bool, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				fn(x)
				output <- true
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x {{.Type}}) {
		input <- x
	})
	close(input)

	Range(0, nrJobs).ForEach(func() {
		<-output
		progressFn()
	})
}

func (rcv {{.TypeName}}List) Count() int {
	return len(rcv)
}

func (rcv {{.TypeName}}List) IsEmpty() bool {
	return len(rcv) == 0
}

func (rcv {{.TypeName}}List) IsNotEmpty() bool {
	return len(rcv) != 0
}

// tests whether this sequence contains the given index
func (rcv {{.TypeName}}List) IsDefinedAt(x int) bool {
	return x >= 0 && x < len(rcv)
}

// returns a slice of {{.Type}}
func (rcv {{.TypeName}}List) ToSlice() []{{.Type}} {
	return rcv
}

// returns the element at the index
func (rcv {{.TypeName}}List) Apply(x int) {{.Type}} {
	return rcv[x]
}

func (rcv {{.TypeName}}List) ApplyOrElse(x int, fn func() {{.Type}}) {{.Type}} {
	if &rcv[x] == nil {
		return rcv[x]
	} else {
		return fn()
	}
}

func (rcv {{.TypeName}}List) FoldLeft(zero {{.Type}}, fn func(acc {{.Type}}, x {{.Type}}) {{.Type}}) {{.Type}} {
	for _, x := range rcv {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv {{.TypeName}}List) FoldRight(zero {{.Type}}, fn func(acc {{.Type}}, x {{.Type}}) {{.Type}}) {{.Type}} {
	for _, x := range rcv.Reverse() {
		zero = fn(zero, x)
	}
	return zero
}

func (rcv {{.TypeName}}List) Intersperse(a {{.Type}}) {{.TypeName}}List {
	ys := Empty{{.TypeName}}List()
	rcv.ForEachWithLastFlag(func(last bool, x {{.Type}}) {
		ys = ys.Append(x)
		if !last {
			ys = ys.Append(a)
		}
	})
	return ys
}

// create a new list prefilled
func (rcv {{.TypeName}}List) Fill(num int, a {{.Type}}) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	for i := 0; i < num; i++ {
		xs = xs.Append(a)
	}
	return xs
}

func (rcv {{.TypeName}}List) Tablulate(num int, fn func(int) {{.Type}}) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	for i := 0; i < num; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv {{.TypeName}}List) Partition(fn func({{.Type}}) bool) ({{.TypeName}}List, {{.TypeName}}List) {
	xs := Empty{{.TypeName}}List()
	ys := Empty{{.TypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		if fn(x) {
			xs = xs.Append(x)
		} else {
			ys = ys.Append(x)
		}
	})
	return xs, ys
}

func (rcv {{.TypeName}}List) MkString(fn func({{.Type}}) string) String {
	var builder strings.Builder
	rcv.ForEach(func(x {{.Type}}) {
		builder.WriteString(fn(x))
	})
	return String(builder.String())
}

func (rcv {{.TypeName}}List) RangeOf(from int, to int, fn func(int) {{.Type}}) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	for i := from; i <= to; i++ {
		xs = xs.Append(fn(i))
	}
	return xs
}

func (rcv {{.TypeName}}List) Contains(a {{.Type}}) bool {
	for _, x := range rcv {
		if cmp.Equal(x, a) {
			return true
		}
	}
	return false
}

func (rcv {{.TypeName}}List) ContainsNot(a {{.Type}}) bool {
	return !rcv.Contains(a)
}

func (rcv {{.TypeName}}List) Distinct() {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		if xs.ContainsNot(x) {
			xs = xs.Append(x)
		}
	})
	return xs
}

// return the intersection of the list and another list
func (rcv {{.TypeName}}List) Intersect(xs {{.TypeName}}List) {{.TypeName}}List {
	ys := Empty{{.TypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		xs.ForEach(func(y {{.Type}}) {
			if cmp.Equal(x, y) && ys.ContainsNot(x) {
				ys = ys.Append(x)
			}
		})
	})
	return ys
}

func (rcv {{.TypeName}}List) Slice(from int, to int) {{.TypeName}}List {
	return rcv[from : to+1]
}

func (rcv {{.TypeName}}List) FlatMapTo{{.TypeName}}List(fn func({{.Type}}) {{.TypeName}}List) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}

func (rcv {{.TypeName}}List) MapTo{{.TypeName}}(fn func({{.Type}}) {{.Type}}) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		xs = xs.Append(fn(x))
	})
	return xs
}

func (rcv {{.TypeName}}List) MapTo{{.TypeName}}WithIndex(fn func(int,{{.Type}}) {{.Type}}) {{.TypeName}}List {
	xs := Empty{{.TypeName}}List()
	rcv.ForEachWithIndex(func(i int, x {{.Type}}) {
		xs = xs.Append(fn(i,x))
	})
	return xs
}

func (rcv {{.TypeName}}List) MapTo{{.TypeName}}P(mapFn func({{.Type}}) {{.Type}}) {{.TypeName}}List {
	return rcv.MapTo{{.TypeName}}PP(10, mapFn)
}

func (rcv {{.TypeName}}List) MapTo{{.TypeName}}PP(parallelism int, mapFn func({{.Type}}) {{.Type}}) {{.TypeName}}List {
	return rcv.MapTo{{.TypeName}}PPP(parallelism, mapFn, func() {})
}

func (rcv {{.TypeName}}List) MapTo{{.TypeName}}PPP(parallelism int, mapFn func({{.Type}}) {{.Type}}, progressFn func()) {{.TypeName}}List {
	nrJobs := rcv.Count()
	input := make(chan {{.Type}}, nrJobs)
	output := make(chan {{.Type}}, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x {{.Type}}) {
		input <- x
	})
	close(input)

	xs := Empty{{.TypeName}}List()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}

// implementation of 'sort.Interface'
func (rcv {{.TypeName}}List) Len() int {
	return rcv.Count()
}

// implementation of 'sort.Interface'
func (rcv {{.TypeName}}List) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

// implementation of sort.Interface
var {{.TypeName}}ListLessFunc = func(i, j int) bool {
	panic("Not implemented")
}

// implementation of sort.Interface
func (rcv {{.TypeName}}List) Less(i, j int) bool {
	return {{.TypeName}}ListLessFunc(i, j)
}

// i and j are two objects that need to be compared, 
// and based on that comparison the List will be sorted
func (rcv {{.TypeName}}List) Sort(fn func(i {{.Type}}, j {{.Type}}) bool) {{.TypeName}}List {
	{{.TypeName}}ListLessFunc = func(i, j int) bool {
		return fn(rcv[i], rcv[j])
	}
	sort.Sort(rcv)
	return rcv
}
`))

var mapToTemplate = template.Must(template.New("generated").Parse(`
func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}(fn func({{.Type}}) {{.ToType}}) {{.ToTypeName}}List {
	ys := make([]{{.ToType}}, 0)
	for _, x := range rcv {
		ys = append(ys, fn(x))
	}
	return ys
}

func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}WithIndex(fn func(int, {{.Type}}) {{.ToType}}) {{.ToTypeName}}List {
	ys := make([]{{.ToType}}, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i, x))
	}
	return ys
}

func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}WithLastFlag(fn func(bool, {{.Type}}) {{.ToType}}) {{.ToTypeName}}List {
	ys := make([]{{.ToType}}, 0)
	for i, x := range rcv {
		ys = append(ys, fn(i+1 == len(rcv), x))
	}
	return ys
}

func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}P(mapFn func({{.Type}}) {{.ToType}}) {{.ToTypeName}}List {
	return rcv.MapTo{{.ToTypeName}}PP(10, mapFn)
}

func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}PP(parallelism int, mapFn func({{.Type}}) {{.ToType}}) {{.ToTypeName}}List {
	return rcv.MapTo{{.ToTypeName}}PPP(parallelism, mapFn, func() {})
}

func (rcv {{.TypeName}}List) MapTo{{.ToTypeName}}PPP(parallelism int, mapFn func({{.Type}}) {{.ToType}}, progressFn func()) {{.ToTypeName}}List {
	nrJobs := rcv.Count()
	input := make(chan {{.Type}}, nrJobs)
	output := make(chan {{.ToType}}, nrJobs)

	// make workers
	Range(0, parallelism).ForEach(func() {
		go func() {
			for x := range input {
				output <- mapFn(x)
			}
		}()
	})

	// put commands on the channel
	rcv.ForEach(func(x {{.Type}}) {
		input <- x
	})
	close(input)

	xs := Empty{{.ToTypeName}}List()
	Range(0, nrJobs).ForEach(func() {
		xs = xs.Append(<-output)
		progressFn()
	})
	return xs
}
`))

var foldMapToTemplate = template.Must(template.New("generated").Parse(`
func (rcv {{.TypeName}}List) FoldMapTo{{.ToTypeName}}(zero {{.ToType}}, foldFn func(acc {{.ToType}}, x {{.ToType}}) {{.ToType}}, mapFn func(x {{.Type}}) {{.ToType}}) {{.ToType}} {
	acc := zero
	for _, x := range rcv {
		y := mapFn(x)
		acc = foldFn(acc, y)
	}
	return acc
}
`))

var importsTemplate = template.Must(template.New("generated").Parse(`
import (
	{{range $index, $import := .Imports}}"{{$import}}"
	{{end}}
)
`))

var intListTemplate = template.Must(template.New("generated").Parse(`
func (rcv IntList) Range(from int, to int) IntList {
	xs := EmptyIntList()
	for i := from; i <= to; i++ {
		xs = xs.Append(i)
	}
	return xs
}
`))

var stringListTemplate = template.Must(template.New("generated").Parse(`
// joins using the character and returns the string
func (rcv StringList) Join(sep string) String {
	return rcv.Intersperse(sep).MkString(func (x string) string { return x})
}
`))

var flatMapToTemplate = template.Must(template.New("generated").Parse(`
func (rcv {{.TypeName}}List) FlatMapTo{{.ToTypeName}}List(fn func({{.Type}}) {{.ToTypeName}}List) {{.ToTypeName}}List {
	xs := Empty{{.ToTypeName}}List()
	rcv.ForEach(func(x {{.Type}}) {
		xs = xs.AppendSlice(fn(x).ToSlice())
	})
	return xs
}
`))
