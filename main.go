package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

type Song struct {
	Title  string
	Artist string
	Album  string
	Year   string
	Lyrics string
}

func isFruuppFile(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close() // auto called when the func returns

	// scans the opened file (for text file)
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		firstLine = strings.TrimSpace(firstLine)

		return firstLine == "#!fruupp"
	}

	return false
}

func listFruuppFiles(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil { return nil, err }

	var fruuppFiles []string
	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := filepath.Join(dirPath, entry.Name())
			if isFruuppFile(filePath) {
				fruuppFiles = append(fruuppFiles, filePath)
			}
		}
	}

	return fruuppFiles, nil
}

func main() {
	isIt := isFruuppFile("./lyrics/get-em-out-by-friday")
	fmt.Printf("%t\n", isIt)
	files, _ := listFruuppFiles("./lyrics") 
	fmt.Println(files)
	fmt.Println(os.Args[1])
}
