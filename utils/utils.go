package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

// ErrorCheck ...
func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// SplitFileByLine ...
func SplitFileByLine(file string) []string {
	data, err := ioutil.ReadFile(file)
	ErrorCheck(err)

	return strings.Split(string(data), "\n")
}
