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
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/template"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
	"sigs.k8s.io/external-dns/pkg/metrics"

	// the imports is necessary for the code generation process.
	_ "sigs.k8s.io/external-dns/controller"
	_ "sigs.k8s.io/external-dns/provider"
	_ "sigs.k8s.io/external-dns/provider/webhook"
)

const markdownTemplate = `# Available Metrics

<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW METRIC ADDED PLEASE RUN 'make generate-metrics-documentation' -->
<!-- markdownlint-disable MD013 -->

All metrics available for scraping are exposed on the {{backtick 1}}/metrics{{backtick 1}} endpoint.
The metrics are in the Prometheus exposition format.

To access the metrics:

{{backtick 3}}sh
curl https://localhost:7979/metrics
{{backtick 3}}

## Supported Metrics

> Full metric name is constructed as follows:
> {{backtick 1}}external_dns_<subsystem>_<name>{{backtick 1}}

| Name                             | Metric Type | Subsystem   |  Help                                                 |
|:---------------------------------|:------------|:------------|:------------------------------------------------------|
{{- range .Metrics }}
| {{ .Name }} | {{ .Type | capitalize }} | {{ .Subsystem }} | {{ .Help }} |
{{- end }}

## Available Go Runtime Metrics

> The following Go runtime metrics are available for scraping. Please note that they may change over time.

| Name                  |
|:----------------------|
{{- range .RuntimeMetrics }}
| {{ . }} |
{{- end -}}
`

func main() {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/docs/monitoring/metrics.md", testPath)
	fmt.Printf("generate file '%s' with configured metrics\n", path)

	content, err := generateMarkdownTable(metrics.RegisterMetric)
	if err != nil {
		_ = fmt.Errorf("failed to generate markdown file '%s': %v\n", path, err.Error())
	}
	content = content + "\n"
	_ = utils.WriteToFile(path, content)
}

func generateMarkdownTable(m *metrics.MetricRegistry) (string, error) {
	tmpl := template.Must(template.New("metrics.md.tpl").Funcs(utils.FuncMap()).Parse(markdownTemplate))

	sortMetrics(m.Metrics)
	runtimeMetrics := getRuntimeMetrics(prometheus.DefaultRegisterer)

	var b bytes.Buffer
	err := tmpl.Execute(&b, struct {
		Metrics        []*metrics.Metric
		RuntimeMetrics []string
	}{
		Metrics:        m.Metrics,
		RuntimeMetrics: runtimeMetrics,
	})

	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// sortMetrics sorts the given slice of metrics by their subsystem and name.
// Metrics are first sorted by their subsystem, and then by their name within each subsystem.
func sortMetrics(metrics []*metrics.Metric) {
	sort.Slice(metrics, func(i, j int) bool {
		if metrics[i].Subsystem == metrics[j].Subsystem {
			return metrics[i].Name < metrics[j].Name
		}
		return metrics[i].Subsystem < metrics[j].Subsystem
	})
}

// getRuntimeMetrics retrieves the list of runtime metrics from the Prometheus library.
func getRuntimeMetrics(reg prometheus.Registerer) []string {
	var runtimeMetrics []string

	// hacks to get the runtime metrics from prometheus library
	// safe to do because it's a just a documentation generator
	values := reflect.ValueOf(reg).Elem().FieldByName("dimHashesByName")
	values = reflect.NewAt(values.Type(), unsafe.Pointer(values.UnsafeAddr())).Elem()

	switch v := values.Interface().(type) {
	case map[string]uint64:
		for k := range v {
			if !strings.HasPrefix(k, "external_dns") {
				runtimeMetrics = append(runtimeMetrics, k)
			}
		}
	}
	sort.Strings(runtimeMetrics)
	return runtimeMetrics
}
