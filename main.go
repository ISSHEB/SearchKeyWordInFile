package main

import (
	"fmt"
	"log"
	"search/pkg"
)

func main() {

	filePaths, err := pkg.GetFilePaths("examples")
	if err != nil {
		log.Fatal(err)
	}
	pkg.Search(filePaths, "Хороши")

	fmt.Print("filePaths", filePaths)
}
