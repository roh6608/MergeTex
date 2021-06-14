package main

import (
	"bufio"
	"flag"
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

func writeFile(filepath string, doc []string) {
	// add error checking
	f, _ := os.Create(filepath)
	defer f.Close()

	for i := 0; i < len(doc); i++ {
		// add error checking
		f.WriteString(doc[i] + "\n")
	}

}

// A function to trim the head and foot of tex files
func getDoc(file []string) []string {
	var start int
	var end int

	for i := 0; i < len(file); i++ {
		if file[i] == `\begin{document}` {
			start = i + 1
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
			merged = append(merged, " ")
		}
	}

	return merged
}

// need to write a function that can create pre-amble for merged document
// need to also add in error checking etc.
// create source and bin directory, also create tests, etc.
// finish writing this version in go then create a new branch for the c version
// function to write the merged document to a file
// function to delete anything that will fail the compilation

func main() {

	// will take command line flag arguments here
	// arguments needed will be, directory of files, merged output name and location
	// later will add options for mergin pre-ambles and formatting etc.
	// add standalone options similar to pandoc, this can be the flag that creates the whole document and doesnt just merge
	// the .tex between begin and end

	inDir := flag.String("i", ".", "The directory containing the files to be merged.")
	out := flag.String("o", ".", "The filename of he merged output.")

	flag.Parse()

	writeFile(*out, merge(*inDir))

}
