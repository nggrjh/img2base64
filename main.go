package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/nggrjh/img2base64/helper"
)

const (
	file = "base64.out"

	typeUrl  = "url"
	typeFile = "file"
)

type obj struct {
	typ  string
	path string
}

var fn = map[string]func(string) []byte{
	typeUrl:  helper.FetchImage,
	typeFile: helper.ReadImage,
}

func main() {
	objs := []obj{
		{typ: typeUrl, path: "https:tekno.esportsku.com/wp-content/uploads/2020/10/Cara-Membuat-KTP-Kucing.jpg"},
	}

	if err := os.RemoveAll(file); err != nil {
		log.Fatal(err)
	}

	for _, o := range objs {
		encoded := encodeUrl(o.path, fn[o.typ])

		// Write the full base64 representation of the image
		write(encoded)
	}

}

func encodeUrl(s string, fn func(string) []byte) string {
	// Convert image to bytes
	bytes := fn(s)

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	var base64Encoding string
	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}

func write(s string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.WriteString(s + "\n"); err != nil {
		log.Fatal(err)
	}
}
