package main

import (
	"bufio"
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

const FruuppShebang string = "#!fruupp"

func parseFruuppFile(filePath string) (*Song, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return  nil, err
	}
	defer file.Close()

	song := &Song{}

	scanner := bufio.NewScanner(file)
	isInMetadata := true
	var lyrics strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == FruuppShebang { continue }
		if strings.TrimSpace(line) == "---" {
			isInMetadata = false
			continue
		}

		// is the scan's current line in metadata section?
		if isInMetadata {
			if strings.Contains(line, ":") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])

					switch strings.ToLower(key) {
					case "title":
						song.Title = value
					case "artist":
						song.Artist = value
					case "year":
						song.Year = value
					case "album":
						song.Album = value
					}
				}
			}
		} else {
			lyrics.WriteString(line)
			lyrics.WriteString("\n")
		}
	}

	// if the scanner returns an error in the end
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	song.Lyrics = lyrics.String()
	return song, nil
}
