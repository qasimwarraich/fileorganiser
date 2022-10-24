package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	files := loadFiles(dir)
	datesmap := createDateMap(files)
	createDatedDirs(dir, datesmap)
	moveFilesByDate(dir, files)
}

func loadFiles(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func createDateMap(files []fs.FileInfo) map[time.Time]int {
	datemap := make(map[time.Time]int)

	for _, file := range files {
		datemap[file.ModTime()] = datemap[file.ModTime()] + 1
	}
	return datemap
}

func createDatedDirs(dir string, datesmap map[time.Time]int) {
	for key := range datesmap {
		date := formatDateString(key)

		_, err := os.Stat(formatDirName(dir, date))
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(formatDirName(dir, date), 0o755)
			}
		}
	}
}

func moveFilesByDate(dir string, files []fs.FileInfo) {
	for _, file := range files {
		if !file.IsDir() {
			proposedPath := formatDateString(file.ModTime())
			proposedPath = formatDirName(proposedPath, file.Name())

			err := os.Rename(formatDirName(dir, file.Name()), formatDirName(dir, proposedPath))
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func formatDirName(parent string, child string) string {
	return parent + "/" + child
}

func formatDateString(modtime time.Time) string {
	return modtime.Format("2006-01-02")
}
