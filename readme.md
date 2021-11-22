# gen
Simple Go app for source code generation

## Motivation
Coming from Haskel/Scala, I miss the abstractions that those collections bring. This small app can be used
to bring a piece of that developer experience back to Golang.

## Effort
Low hanging fruit, high bang-for-buck, quick-win, code generation FTW!

## What does it do?
In the package where it is executed, it will write a `<type_name>_list.go` file that contains well known methods from Scala/Scalaz/Cats without the typeclasses.

## Usage

Either of the CLI or as part of `go:generate`

```
// generate a list 
gen list -p <package_name> -t <type_name> -m <map_to> -f <fold_to>

// generate an option type
gen option -p <package_name> -t <type_name> -m <map_to> -f <fold_to>

// generate a range type
gen range -p <package_name>

// generate a string type
gen string -p <package_name>
```

and/or in your code

```
//go:generate gen list -p package_name -t type_name -f fold_to
```

## Build
Clone the repository and then `go build` or `go get -u github.com/dnvriend/gen` if you just want the gen binary on your path.

Type `which gen` to find out where it is to uninstall it, or `rm $(which gen)` to remove it.

