//go:generate gen string -p main
//go:generate gen range -p main
//go:generate gen list -p main -t int -m string
//go:generate gen option -p main -t int
//go:generate gen list -p main -t string -m DadJoke
//go:generate gen option -p main -t string
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//go:generate gen list -p main -t DadJoke
//go:generate gen option -p main -t DadJoke
type DadJoke struct {
	Id   string `json:"id"`
	Joke string `json:"joke"`
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
	EmptyPersonList().HeadOption().ForEach(func(p Person) {})
	people.HeadOption().ForEach(func(p Person) { fmt.Println("The first person is: ", p) })

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

	// map/reduce to text string
	// make 10 requests
	result := Range(0, 10).
		ToIntList().
		// map to url
		MapToString(func(i int) string {
			return "https://icanhazdadjoke.com/"
		}).
		// make 10 *parallel* requests and wait until all are ready
		// see the 'P', 'PP' and 'PPP' implementations for runner
		// configuration and progress call back functions (eg. progress bar)
		MapToStringP(func(url string) string {
			return get(url)
		}).
		// reduce the result
		MkString(func(s string) string {
			return s
		})
	// print the result
	fmt.Println(result)

	// unmarshal some dad jokes in a pipeline
	listOfDadJoke := Range(0, 30).
		ToIntList().
		MapToString(func(i int) string { return "https://icanhazdadjoke.com/" }).
		MapToStringPP(30, func(url string) string { return get(url) }).
		MapToDadJoke(func(jsonString string) DadJoke {
			joke := DadJoke{}
			json.Unmarshal([]byte(jsonString), &joke)
			return joke
		})

	listOfDadJoke.
		Sort(func(i DadJoke, j DadJoke) bool { return i.Id < j.Id }).
		ForEach(func(joke DadJoke) {
			fmt.Println(joke)
		})
}

func get(url string) string {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
