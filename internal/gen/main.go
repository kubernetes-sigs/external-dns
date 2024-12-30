package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	cfg "sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

type Flag struct {
	Name        string
	Description string
}
type Flags []Flag

// AddFlag adds a new flag to the Flags struct
func (f *Flags) AddFlag(name, description string) {
	*f = append(*f, Flag{Name: name, Description: description})
}

const markdownTemplate = `
# Flags

<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW FLAG ADDED PLEASE RUN go .... -->

| Flag | Description  |
| :------ | :----------- |
{{- range . }}
| {{ .Name }} | {{ .Description }} | {{- end -}}
`

// It generates a markdown file
// with the supported flags and writes it to the 'docs/flags.md' file.
// to re-generate `docs/flags.md` execute 'go run internal/gen/main.go'
func main() {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/docs/flags.md", testPath)
	fmt.Println(fmt.Sprintf("generate file '%s' with supported flags", path))

	flags := computeFlags()
	content, err := flags.generateMarkdownTable()
	if err != nil {
		_ = fmt.Errorf("failed to generate markdown file '%s': %v", path, err.Error())
	}
	_ = writeToFile(path, content)
}

func computeFlags() Flags {
	app := cfg.App(&cfg.Config{})
	modelFlags := app.Model().Flags

	flags := Flags{}

	for _, flag := range modelFlags {
		// do not include helpers and completion flags
		if strings.Contains(flag.Name, "help") || strings.Contains(flag.Name, "completion-") {
			continue
		}
		flagString := ""
		flagName := flag.Name
		if flag.IsBoolFlag() {
			flagName = "[no-]" + flagName
		}
		flagString += fmt.Sprintf("--%s", flagName)

		if !flag.IsBoolFlag() {
			flagString += fmt.Sprintf("=%s", flag.FormatPlaceHolder())
		}
		flags.AddFlag(fmt.Sprintf("`%s`", flagString), flag.HelpWithEnvar())
	}
	return flags
}

func (f *Flags) generateMarkdownTable() (string, error) {
	tmpl := template.Must(template.New("flags.md.tpl").Parse(markdownTemplate))

	var b bytes.Buffer
	err := tmpl.Execute(&b, f)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func writeToFile(filename string, content string) error {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		_ = fmt.Errorf("failed to create file: %v", fileErr)
	}
	defer file.Close()
	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		_ = fmt.Errorf("failed to write to file: %s", filename)
	}
	return nil
}
