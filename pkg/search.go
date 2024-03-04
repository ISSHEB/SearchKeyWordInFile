package pkg

import (
	"bufio"
	"fmt"
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
	keyword = strings.ToLower(keyword)
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("файл не открываеться", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		found := false

		for scanner.Scan() {
			line := strings.ToLower(scanner.Text())
			if strings.Contains(line, keyword) {
				fmt.Println("нашли слово в этом файле:", filePath)
				filepath = append(filepath, filePath)
				found = true
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Неудалось прочитать", err)
		}
		if !found {
			fmt.Println("не найдено")
		}
	}
	return filepath
}
