package pkg

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func GetFilePaths(root string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

func Search(filePaths []string, keyword string) (filepath []string) {
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
				filepath = append(filepath, filePath)
				found = true
				break
			}

		}
		if err := scanner.Err(); err != nil {
			println("Неудалось прочитать", err)
		}
		if !found {
			println("не найдено")
		}
	}
	return filepath
}
