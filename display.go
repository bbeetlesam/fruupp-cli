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
	fmt.Println("just a placeholder. supersister got a present from nancy!")
}

func displayError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
	fmt.Fprintln(os.Stderr, "Try 'fruupp help' to sharpen your dull mind.")
}

func displayVersion() {
	fmt.Println("fruupp v19.7.3")
}
