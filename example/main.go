//go:generate gen list -p main -t int
//go:generate gen option -p main -t int
//go:generate gen list -p main -t string
//go:generate gen option -p main -t string
package main

import "fmt"

//go:generate gen list -p main -t Person
//go:generate gen option -p main -t Person
type Person struct {
	Name      string
	Age       int
	Addresses AddressList
	Cats      CatList
}

//go:generate gen list -p main -t Cat
//go:generate gen option -p main -t Cat
type Cat struct {
	Name string
	Age  int
}

func (rcv Cat) Miauwing() string {
	return "miauw"
}

//go:generate gen list -p main -t Address
//go:generate gen option -p main -t Address
type Address struct {
	Street      string
	HouseNumber int
	Zip         string
}

func main() {
	smith := Person{
		Name: "Mr Smith",
		Age:  47,
		Cats: EmptyCatList().
			Append(Cat{
				Name: "Tiger",
				Age:  12,
			}),
		Addresses: EmptyAddressList().
			Append(Address{
				Street:      "streetname",
				HouseNumber: 123456789,
				Zip:         "123456789",
			}),
	}
	black := Person{
		Name: "Mr Black",
		Age:  21,
		Cats: EmptyCatList().
			Append(Cat{
				Name: "Mouse",
				Age:  1,
			}),
		Addresses: EmptyAddressList().
			Append(Address{
				Street:      "streetname",
				HouseNumber: 123456789,
				Zip:         "123456789",
			}),
	}

	people := EmptyPersonList().
		Append(smith).
		Append(black)

	fmt.Println("people contains black?: ", people.Contains(black))

	fmt.Println("is there a first person?: ", people.HeadOption().IsNotEmpty())

	// it is safe the do this
	EmptyPersonList().HeadOption().ForEach(func (p Person) {})
	people.HeadOption().ForEach(func (p Person) { fmt.Println("The first person is: ", p)})

	people.ForEach(func(p Person) {
		fmt.Println(">Person: ", p.Name, p.Age)
		p.Addresses.ForEach(func(a Address) {
			fmt.Println("* Address: ")
			fmt.Println("  - Street: ", a.Street)
			fmt.Println("  - HouseNumber: ", a.HouseNumber)
			fmt.Println("  - Zip: ", a.Zip)
		})
		p.Cats.ForEach(func(c Cat) {
			fmt.Println("* Cat: ")
			fmt.Println("  - Name: ", c.Name)
			fmt.Println("  - Age: ", c.Age)
		})
	})

	Range(0, 5).ForEach(func() {
		fmt.Println("Foo")
	})
}
