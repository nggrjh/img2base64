package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nggrjh/img2base64/helper"
)

const (
	fileIn  = "in/%s"
	fileOut = "out/%s.out"

	typeUrl  = "url"
	typeFile = "file"

	prefixFlag = false
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
		{typ: typeFile, path: "ktp_ariel.png"},
		{typ: typeFile, path: "ktp_ivan.png"},
		{typ: typeFile, path: "ktp_mira.png"},
		{typ: typeFile, path: "ktp_bintik.png"},
		{typ: typeFile, path: "selfie_ariel.png"},
		{typ: typeFile, path: "selfie_tirto.png"},
		{typ: typeFile, path: "IMG_4885.PNG"},
	}

	if err := os.RemoveAll(fileOut); err != nil {
		log.Fatal(err)
	}

	for _, o := range objs {
		encoded := encodeUrl(fmt.Sprintf(fileIn, o.path), fn[o.typ])

		path := fmt.Sprintf(fileOut, strings.Split(o.path, ".")[0])
		// Write the full base64 representation of the image
		write(path, encoded)
	}

}

func encodeUrl(path string, fn func(string) []byte) string {
	// Convert image to bytes
	bytes := fn(path)

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	var base64Encoding string

	if prefixFlag {
		// Prepend the appropriate URI scheme header depending
		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
		}
	}

	// Append the base64 encoded output
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}

func write(path, text string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.WriteString(text + "\n"); err != nil {
		log.Fatal(err)
	}
}
