package test

//go:generate gen list -p test -t Cat
type Cat struct {
	Name string
	Age  int
}
