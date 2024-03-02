package pkg

import (
	"bufio"
	"os"
	"strings"
)

//type Searcher struct {
//	Files []string `json:"files"`
//	Error string   `json:"error"`
//}

func Search(filePaths []string, keyword string) {
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			println("файл не открываеться", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		found := false

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, keyword) {
				println("нашли слово в этом файле:", filePath)
			}
			if found {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			println("Неудалось прочитать", err)
		}
	}

}
