package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
)

var currentDir string

func main() {
	setCurrentDir()

	files, err := ioutil.ReadDir(path.Join(currentDir, "..", "..", "src"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d files found:\n", len(files))

	for _, file := range files {
		fmt.Printf(" - %s (%d bytes)\n", file.Name(), file.Size())
	}

	var _lines []string

	for _, file := range files {

		file, err := os.Open(path.Join(currentDir, "..", "..", "src", file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() != "" {
				_lines = append(_lines, scanner.Text())
			}
		}
	}

	sort.Slice(_lines, func(i, j int) bool { return _lines[i] < _lines[j] })

	var lines []string
	var duplicates = 0
	for i := 0; i < len(_lines); i++ {
		if i != 0 {
			if lines[i-1-duplicates] == _lines[i] {
				duplicates++
			} else {
				lines = append(lines, _lines[i])
			}
		} else {
			lines = append(lines, _lines[i])
		}
	}

	err = ioutil.WriteFile(path.Join(currentDir, "..", "..", "domains.txt"), []byte(strings.Join(lines, "\n")), 0664)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d total lines\n", len(_lines))
	fmt.Printf("%d unique lines\n", len(lines))
	fmt.Printf("%d duplicate entry\n", duplicates)
	fmt.Printf("%v duplicate entry\n", (float64(duplicates) / float64(len(_lines)) * 100))

}

func setCurrentDir() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir = path.Dir(currentFilePath)
}
