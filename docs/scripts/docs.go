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
