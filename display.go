package main

import (
	"fmt"
	"os"
)

func displayDefaultMsg() {
	fmt.Println("FruuppCLI - A command-line tool for displaying custom song lyrics.")
	fmt.Println("\nType 'fruupp help' for usage.")
}

func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  fruupp <command> [args]")
	fmt.Println("\nCommands:")
	fmt.Println("  whatevershebringwesing")
	fmt.Println("  housewithnodoor")
}

func displayListSongs() {
	var dirPath string
	if len(os.Args) >= 2 && os.Args[1] == "list" {
		if len(os.Args) >= 3 {
			dirPath = os.Args[2]
		} else {
			dirPath = "./"
		}
	}

	files, err := listFruuppFiles(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			displayError("Directory doesn't exist.")
		} else {
			displayError(fmt.Sprintf("Failed to list files: %v\n", err))
		}
		return
	}

	if len(files) == 0 {
		fmt.Println("No fruupp files found.")
		return
	}

	for _, file := range files {
		fmt.Printf("%s\n", file)
	}
}

func displaySong(song *Song) {
	fmt.Printf("Title: %s\n", song.Title)
	fmt.Printf("Artist: %s\n", song.Artist)
	fmt.Printf("Album: %s\n", song.Album)
	fmt.Printf("Released Year: %s\n\n", song.Year)
	fmt.Println(song.Lyrics)
}

func displayError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
	fmt.Fprintln(os.Stderr, "Try 'fruupp help' to sharpen your dull mind.")
}

func displayVersion() {
	fmt.Println("fruupp v19.7.3")
}
