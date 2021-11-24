package buildinfo

import "text/template"

var baseTmpl = template.Must(template.New("generated").Parse(`
type buildInfo struct {
	ShortCommitHash string
	LongCommitHash string
	CurrentDateTime string
}

var BuildInfo = buildInfo {
	ShortCommitHash: "{{.ShortCommitHash}}",
	LongCommitHash: "{{.LongCommitHash}}",
	CurrentDateTime: "{{.CurrentDateTime}}",
}
`))
