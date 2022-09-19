package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := []string{
		"http://i.imgur.com/m1UIjW1.jpg",
	}

	for _, url := range urls {
		encoded := encode(url)

		// Write the full base64 representation of the image
		write(encoded)
	}

}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func encode(url string) string {
	// Fetch image from URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = response.Body.Close() }()

	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)
	return base64Encoding
}

func write(s string) {
	f, err := os.OpenFile("base64.out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	if _, err := f.WriteString(s + "\n"); err != nil {
		log.Fatal(err)
	}
}
