package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var path string
	if len(os.Args) > 1 {

		path = os.Args[1]

	} else {
		path = "./"

	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
