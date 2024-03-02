package main

import "search/pkg"

func main() {
	filePaths := []string{"file1.txt", "file2.txt", "file3.txt"}
	pkg.Search(filePaths, "keyword")
}
