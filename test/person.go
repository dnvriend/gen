package test

//go:generate gen range -p test
//go:generate gen list -p test -t int -f string -f cat -m string
//go:generate gen list -p test -t string -m int -f cat

//go:generate gen list -p test -t Person -m Cat
type Person struct {
	Name      string
	Age       int
	Addresses AddressList
	Cats      CatList
}

//go:generate gen list -p test -t Address
type Address struct {
	Street      string
	HouseNumber int
	Zip         string
}
