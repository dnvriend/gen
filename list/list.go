package list

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dnvriend/gen/collections"
)

func Generate(packageName string, typeName string, mapTo []string, foldMapTo []string, imports []string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("// Generated code; DO NOT EDIT.\npackage %v\n", packageName))
	builder.WriteString(buildImports(imports))
	builder.WriteString(buildBase(packageName, typeName))
	for _, toType := range mapTo {
		builder.WriteString(buildMapTo(typeName, toType))
	}
	for _, toType := range foldMapTo {
		builder.WriteString(buildFoldMapTo(typeName, toType))
	}
	switch typeName {
	case "int":
		builder.WriteString(buildIntListExtras())
	}
	return builder.String()
}

func buildImports(imports []string) string {

	model := struct {
		Imports []string
	}{
		Imports: collections.
			EmptyStringList().
			AppendAll(imports...).
			Append("fmt").
			Append("strings").
			Append("github.com/google/go-cmp/cmp").
			ToSlice(),
	}
	var buf bytes.Buffer
	if err := importsTemplate.Execute(&buf, model); err != nil {
		fmt.Println("generating imports: %v", err)
	}
	return buf.String()
}

func buildMapTo(typeName string, toType string) string {
	model := struct {
		Type            string
		TypeName        string
		ToType          string
		ToTypeName      string
		ToShortTypeName string
	}{
		ToType:          toType,
		ToTypeName:      toType,
		ToShortTypeName: toShortTypeName(toType),
		TypeName:        fixTypeName(typeName),
		Type:            typeName,
	}
	var buf bytes.Buffer
	if err := mapToTemplate.Execute(&buf, model); err != nil {
		fmt.Println("generating map code: %v", err)
	}
	return buf.String()
}

func fixTypeName(name string) string {
	str := strings.ReplaceAll(name, "*", "")
	str = strings.ReplaceAll(str, ".", "")
	return strings.Title(str)
}

func toShortTypeName(name string) string {
	return collections.EmptyStringList().
		AppendSlice(strings.Split(name, ".")).
		Last()
}

func buildFoldMapTo(typeName string, toType string) string {
	model := struct {
		Type       string
		TypeName   string
		ToType     string
		ToTypeName string
	}{
		ToType:     toType,
		ToTypeName: fixTypeName(toType),
		TypeName:   fixTypeName(typeName),
		Type:       typeName,
	}
	var buf bytes.Buffer
	if err := foldMapToTemplate.Execute(&buf, model); err != nil {
		fmt.Println("generating foldmap code: %v", err)
	}
	return buf.String()
}

func buildBase(packageName string, typeName string) string {
	model := struct {
		PackageName string
		TypeName    string
		Type        string
	}{
		PackageName: packageName,
		TypeName:    fixTypeName(typeName),
		Type:        typeName,
	}
	var buf bytes.Buffer
	if err := baseTmpl.Execute(&buf, model); err != nil {
		fmt.Println("generating base code: %v", err)
	}
	return buf.String()
}

func buildIntListExtras() string {
	var buf bytes.Buffer
	if err := intListTemplate.Execute(&buf, nil); err != nil {
		fmt.Println("generating int list extras code: %v", err)
	}
	return buf.String()
}
