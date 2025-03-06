/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
	cfg "sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

type Flag struct {
	Name        string
	Description string
}
type Flags []Flag

// AddFlag adds a new flag to the Flags struct
func (f *Flags) addFlag(name, description string) {
	*f = append(*f, Flag{Name: name, Description: description})
}

// It generates a markdown file
// with the supported flags and writes it to the 'docs/flags.md' file.
// to re-generate `docs/flags.md` execute 'go run internal/gen/docs/flags/main.go'
func main() {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/docs/flags.md", testPath)
	fmt.Printf("generate file '%s' with supported flags\n", path)

	flags := computeFlags()
	content, err := flags.generateMarkdownTable()
	if err != nil {
		_ = fmt.Errorf("failed to generate markdown file '%s': %v", path, err.Error())
	}
	content = content + "\n"
	_ = utils.WriteToFile(path, content)
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
		flags.addFlag(fmt.Sprintf("`%s`", flagString), flag.HelpWithEnvar())
	}
	return flags
}

func (f *Flags) generateMarkdownTable() (string, error) {
	tmpl := template.New("").Funcs(utils.FuncMap())
	template.Must(tmpl.ParseFS(templates, "templates/*.gotpl"))

	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, "flags.gotpl", f)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
