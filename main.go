package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		displayDefaultMsg()
		return
	}

	switch os.Args[1] {
	case "help":
		displayHelp()
	case "version":
		displayVersion()
	case "list":
		displayListSongs()
	case "view":
		if len(os.Args) >= 3 {
			filePath := os.Args[2]

			if !isFruuppFile(filePath) {
				displayError("Not a fruupp file.")
				os.Exit(1)
			}

			song, err := parseFruuppFile(filePath)
			if err != nil {
				if os.IsNotExist(err) {
					// specific error: file doesnt exist
					displayError("File doesn't exist.")
				} else {
					// any other error: perm denied, scanning error, etc
					displayError(err.Error())
				}
				os.Exit(1)
			}

			displaySong(song)
		} else {
			displayError("No file provided.")
			os.Exit(1)
		}
	default:
		displayError("Unknown command.")
		os.Exit(1)
	}
}
