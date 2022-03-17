package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func generateRedirect() {
	version := os.Getenv("EXTERNAL_DNS_VERSION")

	content, err := ioutil.ReadFile("./docs/scripts/index.html.gotmpl")

	if err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir("./docs/redirect", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./docs/redirect/index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("redirect").Parse(string(content)))

	tmpl.Execute(f, version)
}

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
	generateRedirect()
}
