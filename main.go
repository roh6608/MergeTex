package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// A function to return the files in the directory passed to it
func files(directory string) []string {
	dir, _ := ioutil.ReadDir(directory)
	var files []string
	for _, file := range dir {
		files = append(files, file.Name())
	}
	return (files)
}

// A function to read the file into a []string type
func readFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// A function to trim the head and foot of tex files
func getDoc(file []string) {

	// grab everything between \begin{document} and \end{document}, will also have to write a function that can right a new
	// pre-amble for the merged document
}

func main() {

	dir := files(".")
	file := readFile(dir[1])

	for i := 0; i < len(file); i++ {
		fmt.Println(file[i])
	}
}
