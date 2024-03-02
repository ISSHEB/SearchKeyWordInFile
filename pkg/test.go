package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSearchFiles(t *testing.T) {
	// Создаем временный каталог для тестов
	tempDir, err := ioutil.TempDir("", "search_files_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Создаем тестовые файлы
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		path := filepath.Join(tempDir, file)
		err := os.WriteFile(path, []byte("example content"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Выполняем поиск
	query := "example"
	foundFiles, err := searchFiles(query)
	if err != nil {
		t.Fatal(err)
	}

	// Проверяем результаты поиска
	expectedFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	if len(foundFiles) != len(expectedFiles) {
		t.Fatalf("Expected %d files, got %d", len(expectedFiles), len(foundFiles))
	}

	for i, file := range foundFiles {
		if file != expectedFiles[i] {
			t.Errorf("Expected file %s, got %s", expectedFiles[i], file)
		}
	}
}
