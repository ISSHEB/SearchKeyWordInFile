package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"search/pkg"
)

func main() {

	gopath := os.Getenv("GOPATH")

	examplePath := filepath.Join(gopath, "src/examples/")

	files, err := filepath.Glob(filepath.Join(examplePath, "*.txt"))
	if err != nil {
		log.Fatal(err)
	}

	pkg.Search(files, "привет", err)

	fmt.Println("GOPATH:", gopath)
	fmt.Println("GOPATH dddd:", examplePath)
}
