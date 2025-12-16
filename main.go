package main

import (
	"os"
)

type Song struct {
	Title  string
	Artist string
	Album  string
	Year   string
	Lyrics string
}

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
	case "ver":
		displayVersion()
	case "list":
		displayListSongs()
	default:
		displayError("Unknown command.")
		os.Exit(1)
	}
}
