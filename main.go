package main

import (
	"fmt"
	"io/ioutil"
)

func files() []string {
	dir, _ := ioutil.ReadDir("./")
	var files []string
	for _, file := range dir {
		files = append(files, file.Name())
	}
	return (files)
}

func main() {
	fmt.Println(files())
}
