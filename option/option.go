package option

import (
	"bytes"
	"fmt"
	"github.com/dnvriend/gen/util"
	"strings"

	"github.com/dnvriend/gen/typ"
)

func Generate(packageName string, typeName string, mapTo []string, foldMapTo []string, imports []string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("// Generated code; DO NOT EDIT.\npackage %v\n", packageName))
	builder.WriteString(buildImports(imports))
	builder.WriteString(buildBase(packageName, typeName))
	return builder.String()
}

func buildImports(imports []string) string {

	model := struct {
		Imports []string
	}{
		Imports: typ.
			EmptyStringList().
			AppendAll(imports...).
			//Append("fmt").
			//Append("strings").
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
		PackageName string
		TypeName    string
		Type        string
	}{
		PackageName: packageName,
		TypeName:    util.TypeName(typeName),
		Type:        typeName,
	}
	var buf bytes.Buffer
	if err := baseTmpl.Execute(&buf, model); err != nil {
		fmt.Printf("generating base code: %v\n", err)
	}
	return buf.String()
}
