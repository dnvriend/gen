package buildinfo

import (
	"bytes"
	"fmt"
	"strings"
)

func Generate(shortCommitHash string, longCommitHash string, currentDateTime string, packageName string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("// Generated code; DO NOT EDIT.\npackage %v\n", packageName))
	builder.WriteString(buildBase(shortCommitHash, longCommitHash, currentDateTime, packageName))
	return builder.String()
}

func buildBase(shortCommitHash string, longCommitHash string, currentDateTime string, packageName string) string {
	model := struct {
		ShortCommitHash string
		LongCommitHash  string
		CurrentDateTime string
		PackageName     string
	}{
		ShortCommitHash: strings.TrimSpace(shortCommitHash),
		LongCommitHash:  strings.TrimSpace(longCommitHash),
		CurrentDateTime: strings.TrimSpace(currentDateTime),
		PackageName:     packageName,
	}
	var buf bytes.Buffer
	if err := baseTmpl.Execute(&buf, model); err != nil {
		fmt.Printf("generating base code: %v\n", err)
	}
	return buf.String()
}
