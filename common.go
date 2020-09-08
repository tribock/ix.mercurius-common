package common

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func appendToDoku(envParam string) {
	err := writeLines(envParam, "./env.md")
	if err != nil {
		panic(err)
	}
}
func fileContainsText(textToAppend string, path string) bool {
	file, err := ioutil.ReadFile(path)
	FailOnError(err, "not able to read File")
	return strings.Contains(string(file), textToAppend)
}
func writeLines(textToAppend string, path string) error {
	if fileContainsText(textToAppend, path) {
		return nil
	}
	// overwrite file if it exists
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// new writer w/ default 4096 buffer size
	w := bufio.NewWriter(file)
	_, err = w.WriteString(textToAppend + "\n")
	if err != nil {
		return err
	}
	// flush outstanding data
	return w.Flush()
}
func Getenv(key, fallback string) string {
	appendToDoku(key)
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
