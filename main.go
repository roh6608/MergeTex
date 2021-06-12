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
func getDoc(file []string) []string {
	var start int
	var end int

	for i := 0; i < len(file); i++ {
		if file[i] == `\begin{document}` {
			start = i
		} else if file[i] == `\end{document}` {
			end = i
		}
	}

	return (file[start:end])
}

func merge(dir string) []string {
	files := files(dir)
	var merged []string
	var file []string

	for i := 0; i < len(files); i++ {
		file = getDoc(readFile(dir + "/" + files[i]))
		for j := 0; j < len(file); j++ {
			merged = append(merged, file[j])
		}
	}

	return merged
}

// need to write a function that can create pre-amble for merged document

func main() {

	// will take command line flag arguments here
}
