// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package view

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Exporter exports the collected records as view data.
//
// The ExportView method should return quickly; if an
// Exporter takes a significant amount of time to
// process a Data, that work should be done on another goroutine.
//
// It is safe to assume that ExportView will not be called concurrently from
// multiple goroutines.
//
// The Data should not be modified.
type Exporter interface {
	ExportView(viewData *Data)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you no longer
// want data to be exported, invoke UnregisterExporter
// with the previously registered exporter.
//
// Binaries can register exporters, libraries shouldn't register exporters.
func RegisterExporter(e Exporter) {
	defaultWorker.RegisterExporter(e)
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
	defaultWorker.UnregisterExporter(e)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

||||||| parent of 5ce8c7613 (update vendored files)
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

=======
>>>>>>> 5ce8c7613 (update vendored files)
// Exporter exports the collected records as view data.
//
// The ExportView method should return quickly; if an
// Exporter takes a significant amount of time to
// process a Data, that work should be done on another goroutine.
//
// It is safe to assume that ExportView will not be called concurrently from
// multiple goroutines.
//
// The Data should not be modified.
type Exporter interface {
	ExportView(viewData *Data)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you no longer
// want data to be exported, invoke UnregisterExporter
// with the previously registered exporter.
//
// Binaries can register exporters, libraries shouldn't register exporters.
func RegisterExporter(e Exporter) {
	defaultWorker.RegisterExporter(e)
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
<<<<<<< HEAD
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
=======
	defaultWorker.UnregisterExporter(e)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

||||||| parent of 6b7ce455e (update vendored files)
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

=======
>>>>>>> 6b7ce455e (update vendored files)
// Exporter exports the collected records as view data.
//
// The ExportView method should return quickly; if an
// Exporter takes a significant amount of time to
// process a Data, that work should be done on another goroutine.
//
// It is safe to assume that ExportView will not be called concurrently from
// multiple goroutines.
//
// The Data should not be modified.
type Exporter interface {
	ExportView(viewData *Data)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you no longer
// want data to be exported, invoke UnregisterExporter
// with the previously registered exporter.
//
// Binaries can register exporters, libraries shouldn't register exporters.
func RegisterExporter(e Exporter) {
	defaultWorker.RegisterExporter(e)
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
<<<<<<< HEAD
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
=======
	defaultWorker.UnregisterExporter(e)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

||||||| parent of 4d7e5ad26 (update vendored files)
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
// Exporter exports the collected records as view data.
//
// The ExportView method should return quickly; if an
// Exporter takes a significant amount of time to
// process a Data, that work should be done on another goroutine.
//
// It is safe to assume that ExportView will not be called concurrently from
// multiple goroutines.
//
// The Data should not be modified.
type Exporter interface {
	ExportView(viewData *Data)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you no longer
// want data to be exported, invoke UnregisterExporter
// with the previously registered exporter.
//
// Binaries can register exporters, libraries shouldn't register exporters.
func RegisterExporter(e Exporter) {
	defaultWorker.RegisterExporter(e)
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
<<<<<<< HEAD
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
=======
	defaultWorker.UnregisterExporter(e)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "sync"

var (
	exportersMu sync.RWMutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

// Exporter exports the collected records as view data.
//
// The ExportView method should return quickly; if an
// Exporter takes a significant amount of time to
// process a Data, that work should be done on another goroutine.
//
// It is safe to assume that ExportView will not be called concurrently from
// multiple goroutines.
//
// The Data should not be modified.
type Exporter interface {
	ExportView(viewData *Data)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you no longer
// want data to be exported, invoke UnregisterExporter
// with the previously registered exporter.
//
// Binaries can register exporters, libraries shouldn't register exporters.
func RegisterExporter(e Exporter) {
	exportersMu.Lock()
	defer exportersMu.Unlock()

	exporters[e] = struct{}{}
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
