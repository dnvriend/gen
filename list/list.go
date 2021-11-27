package list

import (
	"bytes"
	"fmt"
	"github.com/dnvriend/gen/typ"
	"github.com/dnvriend/gen/util"
)

func Generate(
	packageName string,
	typeName string,
	mapTo typ.StringList,
	flatMapTo typ.StringList,
	foldMapTo typ.StringList,
	imports typ.StringList) string {

	ws := mapTo.MapToStringList(func(toType string) string {
		return buildMapTo(typeName, toType)
	})

	xs := flatMapTo.MapToStringList(func(toType string) string {
		return buildFlatMapTo(typeName, toType)
	})

	ys := foldMapTo.MapToStringList(func(toType string) string {
		return buildFoldMapTo(typeName, toType)
	})

	zs := typ.EmptyStringList().
		Append(fmt.Sprintf("// Generated code; DO NOT EDIT.\npackage %v\n", packageName)).
		Append(buildImports(imports)).
		Append(buildBase(packageName, typeName)).
		AppendSlice(ws.ToSlice()).
		AppendSlice(xs.ToSlice()).
		AppendSlice(ys.ToSlice())

	switch typeName {
	case "int":
		zs = zs.Append(buildIntListExtras())
	case "string":
		zs = zs.Append(buildStringListExtras())
	}

	return zs.MkString()
}

func buildImports(imports []string) string {

	model := struct {
		Imports []string
	}{
		Imports: typ.
			EmptyStringList().
			AppendAll(imports...).
			Append("sort").
			Append("strings").
			Append("github.com/google/go-cmp/cmp").
			ToSlice(),
	}
	var buf bytes.Buffer
	if err := importsTemplate.Execute(&buf, model); err != nil {
		fmt.Printf("generating imports: %v\n", err)
	}
	return buf.String()
}

func buildBase(packageName string, typeName string) string {
	model := struct {
		Type        string // the original type
		TypeName    string // [T] of the container type, based on the type
		PackageName string // the name of the package
	}{
		Type:        typeName,
		TypeName:    util.TypeName(typeName),
		PackageName: packageName,
	}
	var buf bytes.Buffer
	if err := baseTmpl.Execute(&buf, model); err != nil {
		fmt.Printf("generating base code: %v\n", err)
	}
	return buf.String()
}

func buildMapTo(typeName string, toType string) string {
	model := struct {
		Type       string
		TypeName   string
		ToType     string
		ToTypeName string
	}{
		Type:       typeName,
		TypeName:   util.TypeName(typeName),
		ToType:     toType,
		ToTypeName: util.TypeName(toType),
	}
	var buf bytes.Buffer
	if err := mapToTemplate.Execute(&buf, model); err != nil {
		fmt.Printf("generating map code: %v\n", err)
	}
	return buf.String()
}

func buildFlatMapTo(typeName string, toType string) string {
	model := struct {
		Type       string
		TypeName   string
		ToType     string
		ToTypeName string
	}{
		Type:       typeName,
		TypeName:   util.TypeName(typeName),
		ToType:     toType,
		ToTypeName: util.TypeName(toType),
	}
	var buf bytes.Buffer
	if err := flatMapToTemplate.Execute(&buf, model); err != nil {
		fmt.Printf("generating flatmap code: %v\n", err)
	}
	return buf.String()
}

func buildFoldMapTo(typeName string, toType string) string {
	model := struct {
		Type       string
		TypeName   string
		ToType     string
		ToTypeName string
	}{
		Type:       typeName,
		TypeName:   util.TypeName(typeName),
		ToType:     toType,
		ToTypeName: util.TypeName(toType),
	}
	var buf bytes.Buffer
	if err := foldMapToTemplate.Execute(&buf, model); err != nil {
		fmt.Printf("generating foldmap code: %v\n", err)
	}
	return buf.String()
}

func buildIntListExtras() string {
	var buf bytes.Buffer
	if err := intListTemplate.Execute(&buf, nil); err != nil {
		fmt.Printf("generating int list extras code: %v\n", err)
	}
	return buf.String()
}

func buildStringListExtras() string {
	var buf bytes.Buffer
	if err := stringListTemplate.Execute(&buf, nil); err != nil {
		fmt.Printf("generating string list extras code: %v\n", err)
	}
	return buf.String()
}
