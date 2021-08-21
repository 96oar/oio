package oio

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func IsExist(dir string) bool {
	exist := false
	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		exist = true
	}
	return exist
}

func CreateDir(dir string) {
	if !IsExist(dir) {
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func CreateFile(pathFile string, data string) {
	err := ioutil.WriteFile(pathFile, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func ReadFile(template string) string {
	data, _ := ioutil.ReadFile(template)
	return string(data)
}

func ReplaceTextInFile(template string, mapForReplace map[string]string) string {
	input := ReadFile(template)

	for key, value := range mapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

func NewFileForTemplate(pathFile string, template string, mapForReplace map[string]string) {
	data := ReplaceTextInFile(template, mapForReplace)
	CreateFile(pathFile, data)
}
