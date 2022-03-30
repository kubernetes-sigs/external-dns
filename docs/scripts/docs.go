/*
Copyright 2017 The Kubernetes Authors.

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
	"log"
	"os"
	"strings"
)

func removeLinkPrefixInIndex() {
	content, err := os.ReadFile("./docs/index.md")
	if err != nil {
		log.Fatalf("Could not read index.md file. Make sure to run copy_docs.sh script first. Original error: %s", err)
	}

	updatedContent := strings.ReplaceAll(string(content), "](./docs/", "](")
	updatedContent = strings.ReplaceAll(updatedContent, "](docs/", "](")

	f, err := os.OpenFile("./docs/index.md", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Could not open index.md file to update content. Original error: %s", err)
	}
	defer f.Close()

	if _, err := f.WriteString(updatedContent); err != nil {
		log.Fatalf("Failed writing links update to index.md. Original error: %s", err)
	}
}

func main() {
	removeLinkPrefixInIndex()
}
