package helper

import (
	"io/ioutil"
	"log"
)

func ReadImage(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
