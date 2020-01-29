package main

import (
    "fmt"
    "os"
    "strings"
)

const (
	TEXT_FILE="TEXT_FILE"
	 TARGET_URL="TARGET_URL"
)

func main() {
	var textFile, targetUrl string;


	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		
		if pair[0] == TEXT_FILE {
			textFile = pair[1]
		} else if pair[0] == TARGET_URL {
			targetUrl = pair[1]
		} 
    }

	var ok bool

	if len(textFile) == 0 {
		ok = false
		fmt.Println("Please provide a TEXT_FILE env var for the location of the file.")
	} else if len(targetUrl) == 0 {
		ok = false
		fmt.Println("Please provide a TARGET_URL env var for the location of the file.")
	} else {
		ok = true
	}

	if !ok {
		panic("Please check env vars")
	}

	fmt.Println("TEXT_FILE=", textFile)
	fmt.Println("TARGET_URL=", targetUrl)
}