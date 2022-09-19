package helper

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchImage(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = response.Body.Close() }()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
