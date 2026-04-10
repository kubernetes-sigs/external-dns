// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

import (
	"fmt"
	"strings"
)

<<<<<<< HEAD
var allModes = []LoadMode{
	NeedName,
	NeedFiles,
	NeedCompiledGoFiles,
	NeedImports,
	NeedDeps,
<<<<<<< HEAD
<<<<<<< HEAD
	NeedExportFile,
	NeedTypes,
	NeedSyntax,
	NeedTypesInfo,
	NeedTypesSizes,
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
var allModes = []LoadMode{
	NeedName,
	NeedFiles,
	NeedCompiledGoFiles,
	NeedImports,
	NeedDeps,
	NeedExportFile,
	NeedTypes,
	NeedSyntax,
	NeedTypesInfo,
	NeedTypesSizes,
=======
var modes = [...]struct {
	mode LoadMode
	name string
}{
	{NeedName, "NeedName"},
	{NeedFiles, "NeedFiles"},
	{NeedCompiledGoFiles, "NeedCompiledGoFiles"},
	{NeedImports, "NeedImports"},
	{NeedDeps, "NeedDeps"},
	{NeedExportFile, "NeedExportFile"},
	{NeedTypes, "NeedTypes"},
	{NeedSyntax, "NeedSyntax"},
	{NeedTypesInfo, "NeedTypesInfo"},
	{NeedTypesSizes, "NeedTypesSizes"},
	{NeedForTest, "NeedForTest"},
	{NeedModule, "NeedModule"},
	{NeedEmbedFiles, "NeedEmbedFiles"},
	{NeedEmbedPatterns, "NeedEmbedPatterns"},
<<<<<<< HEAD
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
=======
	{NeedTarget, "NeedTarget"},
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
}

<<<<<<< HEAD
var modeStrings = []string{
	"NeedName",
	"NeedFiles",
	"NeedCompiledGoFiles",
	"NeedImports",
	"NeedDeps",
	"NeedExportFile",
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	NeedExportsFile,
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	NeedExportsFile,
=======
	NeedExportFile,
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	NeedTypes,
	NeedSyntax,
	NeedTypesInfo,
	NeedTypesSizes,
}

var modeStrings = []string{
	"NeedName",
	"NeedFiles",
	"NeedCompiledGoFiles",
	"NeedImports",
	"NeedDeps",
<<<<<<< HEAD
	"NeedExportsFile",
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"NeedExportsFile",
=======
	"NeedExportFile",
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"NeedTypes",
	"NeedSyntax",
	"NeedTypesInfo",
	"NeedTypesSizes",
}

func (mod LoadMode) String() string {
	m := mod
	if m == 0 {
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
var modeStrings = []string{
	"NeedName",
	"NeedFiles",
	"NeedCompiledGoFiles",
	"NeedImports",
	"NeedDeps",
	"NeedExportFile",
	"NeedTypes",
	"NeedSyntax",
	"NeedTypesInfo",
	"NeedTypesSizes",
}

func (mod LoadMode) String() string {
	m := mod
	if m == 0 {
=======
func (mode LoadMode) String() string {
	if mode == 0 {
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
		return "LoadMode(0)"
	}
	var out []string
	// named bits
	for _, item := range modes {
		if (mode & item.mode) != 0 {
			mode ^= item.mode
			out = append(out, item.name)
		}
	}
	// unnamed residue
	if mode != 0 {
		if out == nil {
			return fmt.Sprintf("LoadMode(%#x)", int(mode))
		}
		out = append(out, fmt.Sprintf("%#x", int(mode)))
	}
	if len(out) == 1 {
		return out[0]
	}
	return "(" + strings.Join(out, "|") + ")"
}
