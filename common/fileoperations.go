package common

import (
	"io/ioutil"
	"log"
)

func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text
}
