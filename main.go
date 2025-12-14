package main

import (
	"fmt"
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
	isIt := isFruuppFile("./lyrics/get-em-out-by-friday")
	fmt.Printf("%t\n", isIt)
	files, _ := listFruuppFiles("./lyrics") 
	fmt.Println(files)
	fmt.Println(os.Args[0])
}
