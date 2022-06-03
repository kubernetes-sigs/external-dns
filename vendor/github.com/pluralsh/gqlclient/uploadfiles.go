package gqlclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
	"github.com/schollz/progressbar/v3"
)

// Upload File represents a file to upload.
type Upload struct {
	Field string
	Name  string
	R     io.Reader
}

func WithFiles(files []Upload) func(req *http.Request) {
	return func(req *http.Request) {
		bd := client.Request{}
		err := json.NewDecoder(req.Body).Decode(&bd)
		if err != nil {
			fmt.Printf("decode body %v", err)
		}
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		if err := bodyWriter.WriteField("query", bd.Query); err != nil {
			fmt.Printf("write query field %v", err)
		}

		var variablesBuf bytes.Buffer
		if len(bd.Variables) > 0 {
			variablesField, err := bodyWriter.CreateFormField("variables")
			if err != nil {
				fmt.Printf("create variables field %v", err)
			}
			if err := json.NewEncoder(io.MultiWriter(variablesField, &variablesBuf)).Encode(bd.Variables); err != nil {
				fmt.Printf("encode variables %v", err)
			}
		}

		for _, f := range files {
			part, err := bodyWriter.CreateFormFile(f.Field, f.Name)
			if err != nil {
				fmt.Printf("create form file %v", err)
			}

			if _, err := io.Copy(part, f.R); err != nil {
				fmt.Printf("preparing file %v", err)
			}
		}
		bodyWriter.Close()

		bar := progressbar.DefaultBytes(
			int64(bodyBuf.Len()),
			"upload progress",
		)
		reader := progressbar.NewReader(bodyBuf, bar)
		defer reader.Close()

		req.ContentLength = int64(bodyBuf.Len())
		req.Body = &reader
		req.Header.Set("Accept", "application/json; charset=utf-8")
		req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	}
}
