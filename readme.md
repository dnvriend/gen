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
gen list -p <package_name> -t <type_name> -m <map_to> -f <fold_to>
```

and/or in your code

```
//go:generate gen list -p package_name -t type_name -f fold_to
```

## Build

```
go build
cp go /usr/local/bin
```
