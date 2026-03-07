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
	"embed"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
	"sigs.k8s.io/external-dns/pkg/metrics"

	// these imports are necessary for the code generation process.
	_ "sigs.k8s.io/external-dns/controller"
	_ "sigs.k8s.io/external-dns/provider"
	_ "sigs.k8s.io/external-dns/provider/webhook"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/docs/monitoring/metrics.md", testPath)
	fmt.Printf("generate file '%s' with configured metrics\n", path)

	content, err := generateMarkdownTable(metrics.RegisterMetric, true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to generate markdown file '%s': %v\n", path, err)
		os.Exit(1)
	}
	content += "\n"
	_ = utils.WriteToFile(path, content)
}

func generateMarkdownTable(m *metrics.MetricRegistry, withRuntime bool) (string, error) {
	sortMetrics(m.Metrics)
	var runtimeMetrics []string
	if withRuntime {
		runtimeMetrics = getRuntimeMetrics(prometheus.DefaultGatherer)
		// available when promhttp.Handler() is activated
		runtimeMetrics = append(runtimeMetrics, []string{
			"process_network_receive_bytes_total",
			"process_network_transmit_bytes_total",
		}...)
		sort.Strings(runtimeMetrics)
		runtimeMetrics = slices.Compact(runtimeMetrics)
	} else {
		runtimeMetrics = []string{}
	}

	return utils.RenderTemplate(templates, "metrics.gotpl", struct {
		Metrics        []*metrics.Metric
		RuntimeMetrics []string
	}{
		Metrics:        m.Metrics,
		RuntimeMetrics: runtimeMetrics,
	})
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

// getRuntimeMetrics retrieves the list of runtime metrics from the Prometheus registry.
func getRuntimeMetrics(gatherer prometheus.Gatherer) []string {
	mfs, err := gatherer.Gather()
	if err != nil {
		return nil
	}

	var runtimeMetrics []string
	for _, mf := range mfs {
		name := mf.GetName()
		if !strings.HasPrefix(name, "external_dns") {
			runtimeMetrics = append(runtimeMetrics, name)
		}
	}
	sort.Strings(runtimeMetrics)
	return runtimeMetrics
}
