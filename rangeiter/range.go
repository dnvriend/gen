package rangeiter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dnvriend/gen/typ"
)

func Generate(packageName string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("// Generated code; DO NOT EDIT.\npackage %v\n", packageName))
	//builder.WriteString(buildImports())
	builder.WriteString(buildBase(packageName))
	return builder.String()
}

func buildImports() string {
	model := struct {
		Imports []string
	}{
		Imports: typ.
			EmptyStringList().
			Append("net/url").
			Append("strings").
			//Append("github.com/google/go-cmp/cmp").
			ToSlice(),
	}
	var buf bytes.Buffer
	if err := importsTemplate.Execute(&buf, model); err != nil {
		fmt.Printf("generating imports: %v\n", err)
	}
	return buf.String()
}

func buildBase(packageName string) string {
	model := struct {
		PackageName string
	}{
		PackageName: packageName,
	}
	var buf bytes.Buffer
	if err := baseTmpl.Execute(&buf, model); err != nil {
		fmt.Printf("generating base code: %v\n", err)
	}
	return buf.String()
}
