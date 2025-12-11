package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	isIt := isFruuppFile("./lyrics/get-em-out-by-friday")
	fmt.Printf("%t\n", isIt)
	fmt.Printf("%t", isFruuppFile("thus, spoke zarathustra"))
}
