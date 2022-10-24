package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	path, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	files := loadFiles(path)
	moveFiles(path, files)
}

func loadFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func createDatedDir(path string, file fs.FileInfo) {
	dateStamp := formatDateString(file.ModTime())
	err := os.MkdirAll(filepath.Join(path, dateStamp), 0o755)
	if err != nil {
		log.Fatal(err)
	}
}

func moveFiles(path string, files []fs.FileInfo) {
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(path, file.Name())

			dateStamp := formatDateString(file.ModTime())
			proposedPath := filepath.Join(path, dateStamp, file.Name())

			createDatedDir(path, file)

			err := os.Rename(filePath, proposedPath)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func formatDateString(modtime time.Time) string {
	return modtime.Format("2006-01-02")
}
