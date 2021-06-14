package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// A function remove duplicates in an string arrays
func removeDup(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

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

// A function to write the merged function to a file
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
			end = i - 1
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

// A function to get the preambles from the .tex files to be merged.
func getPreamble(file []string) []string {
	var start int
	var end int

	for i := 0; i < len(file); i++ {
		if file[i][0:13] == `\documentclass` {
			start = i + 1
		} else if file[i] == `\begin{document}` {
			end = i - 1
		} else {
			fmt.Println(`The .tex file is malformed, it is missing \documentclass and/or \begin{document}.`)

		}

	}

	preamble := file[start:end]

	return preamble

}

// A function to merge the preambles of the .tex files
func mergePreamble(dir string) []string {
	files := files(dir)
	var merged []string
	var preamble []string

	for i := 0; i < len(dir); i++ {
		preamble = getPreamble(readFile(dir + "/" + files[i]))
		for j := 0; j < len(preamble); j++ {
			merged = append(merged, preamble[j])
		}
	}

	return merged

}

// A function to return a standalone document that should be ready to be compiled
func standaloneDoc(dir string) []string {
	var documentclass []string
	var begin []string
	var end []string
	documentclass[0] = `\documentclass[12pt]{article}`
	begin[0] = `\begin{document}`
	end[0] = `\end{document}`

	tmp := append(documentclass, mergePreamble(dir)...)
	tmp = append(tmp, begin...)
	return append(tmp, end...)

}

func main() {

	// allow an option where they can pass a pre-amble as well
	// later will add options for mergin pre-ambles and formatting etc.
	// -s flag for standalone merged tex file

	in := flag.String("i", ".", "The directory containing the files to be merged.")
	out := flag.String("o", ".", "The filename of he merged output.")
	standalone := flag.Bool("s", false, "If the merged document should be standalone")

	flag.Parse()

	if *standalone == false {
		writeFile(*out, merge(*in))

	} else if *standalone == true {
		writeFile(*out, standaloneDoc(*in))
	}

}
