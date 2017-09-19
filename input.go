package cryptopals

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

func InputBytes(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return content
}

func InputString(path string) string {
	return strings.TrimSpace(string(InputBytes(path)))
}

func InputLines(path string) []string {
	return strings.Split(InputString(path), "\n")
}

func InputScanner(path string) *bufio.Scanner {
	return bufio.NewScanner(bytes.NewReader(InputBytes(path)))
}
