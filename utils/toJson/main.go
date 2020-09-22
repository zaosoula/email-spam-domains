package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

//Usage go run filename -text=dataYouAreLookingfor
//if looking for Nissan in file the command will be
// go run filename -text=Nissan
var currentDir string

func main() {
	setCurrentDir()

	file, err := os.Open(path.Join(currentDir, "..", "..", "domains.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	domains := []string{}
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	jsonDomains, _ := json.Marshal(domains)
	fmt.Println(string(jsonDomains))

	err = ioutil.WriteFile(path.Join(currentDir, "..", "..", "domains.json"), []byte(string(jsonDomains)), 0664)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(domains)

}

func setCurrentDir() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir = path.Dir(currentFilePath)
}
