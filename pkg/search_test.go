package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilePaths(t *testing.T) {
	dir := createTestDir(t)
	defer os.RemoveAll(dir)
	root := "testdata"
	filePaths, err := GetFilePaths(root)
	if err != nil {
		t.Fatalf("ошибка при получении путей к файлам: %v", err)
	}
	if len(filePaths) != 2 {
		t.Fatalf("ожидалось 2 файла, но найдено %d", len(filePaths))
	}
	expectedFilePaths := []string{
		filepath.Join(root, "file1.txt"),
		filepath.Join(root, "file2.txt"),
	}
	for i, filePath := range filePaths {
		if filePath != expectedFilePaths[i] {
			t.Errorf("ожидалось %s, но найдено %s", expectedFilePaths[i], filePath)
		}
	}
}

func TestSearch(t *testing.T) {
	dir := createTestDir(t)
	defer os.RemoveAll(dir)
	root := "testdata"
	filePaths, err := GetFilePaths(root)
	if err != nil {
		t.Fatalf("ошибка при получении путей к файлам: %v", err)
	}
	keyword := "word1"
	foundFilePaths := Search(filePaths, keyword)
	if len(foundFilePaths) != 1 {
		t.Fatalf("ожидалось 1 файл с ключевым словом, но найдено %d", len(foundFilePaths))
	}
	expectedFilePath := filepath.Join(root, "file1.txt")
	if foundFilePaths[0] != expectedFilePath {
		t.Errorf("ожидалось %s, но найдено %s", expectedFilePath, foundFilePaths[0])
	}
}
func createTestDir(t *testing.T) string {
	dir, err := os.MkdirTemp("", "testdata")
	if err != nil {
		t.Fatalf("ошибка при создании временного каталога: %v", err)
	}
	return dir
}
