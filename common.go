package common

import (
	"bufio"
	"fmt"
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
func fileContainsText(textToAppend string, file *os.File) bool {
	scanner := bufio.NewScanner(file)
	b, err := ioutil.ReadAll(file)
	FailOnError(err, "not able to read File")
	fmt.Print(b)
	log.Println(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		log.Println("HODOOOOOR")
		log.Println(scanner.Text(), textToAppend)
		if strings.Contains(scanner.Text(), textToAppend) {
			return true
		}
	}
	return false
}
func writeLines(textToAppend string, path string) error {
	// overwrite file if it exists
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	log.Println(textToAppend)
	if fileContainsText(textToAppend, file) {
		return err
	}
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
