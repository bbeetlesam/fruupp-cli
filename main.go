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
		printDefaultMsg()
	} else if os.Args[1] == "help" {
		printHelp()
	} else if os.Args[1] == "list" {
		printList()
	} else if os.Args[1] == "version" || os.Args[1] == "ver" {
		printVersion()
	} else {
		printError("Unknown command.")
	}
}
