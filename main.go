package main

import (
	"log"
	"search/pkg"
)

func main() {

	filePaths, err := pkg.GetFilePaths("examples")
	if err != nil {
		log.Fatal(err)
	}
	pkg.Search(filePaths, "Хороши")

	//searchFlag := flag.String("search", "", "Keyword to search for")
	//filePathArg := flag.String("file", "", "File to search in")
	//
	//flag.Parse()
	//
	//if *searchFlag == "" {
	//	fmt.Println("Please provide a keyword using the -search flag")
	//	os.Exit(1)
	//}
	//
	//if *filePathArg == "" {
	//	fmt.Println("Please provide a file path using the -file flag")
	//	os.Exit(1)
	//}
	//
	//filePaths := []string{*filePathArg}
	//Search(filePaths, *searchFlag)

}
