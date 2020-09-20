package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

//Usage go run filename -text=dataYouAreLookingfor
//if looking for Nissan in file the command will be
// go run filename -text=Nissan
var currentDir string

func main() {
	setCurrentDir()

	var text string
	// use it as cmdline argument
	textArg := flag.String("domain", "", "Domain to check for")
	flag.Parse()
	// if cmdline arg was not passed ask
	if fmt.Sprintf("%s", *textArg) == "" {
		flag.PrintDefaults()
	} else {
		text = fmt.Sprintf("%s", *textArg)

		file, err := os.Open(path.Join(currentDir, "..", "..", "domains.txt"))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), text) {
				fmt.Println(1)
				return
			}
		}

		fmt.Println(0)
	}

}

func setCurrentDir() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir = path.Dir(currentFilePath)
}
