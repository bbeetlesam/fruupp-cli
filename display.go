package main

import (
	"fmt"
	"os"
)

func printDefaultMsg() {
	fmt.Println("FruuppCLI - A command-line tool for displaying custom song lyrics.")
	fmt.Println("\nType 'fruupp help' for usage.")
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  fruupp <command> [args]")
	fmt.Println("\nCommands:")
	fmt.Println("  whatevershebringwesing")
	fmt.Println("  housewithnodoor")
}

func printList() {
	fmt.Println("just a placeholder. supersister got a present from nancy!")
}

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
	fmt.Fprintln(os.Stderr, "Try 'fruupp help' to sharpen your dull mind.")
}

func printVersion() {
	fmt.Println("fruupp v1973")
}
