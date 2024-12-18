package main

import (
	"fmt"
	filecopier "go_projects/file-manager-projects/file_copier"
	textreverser "go_projects/file-manager-projects/text_reverser"
	wordcounter "go_projects/file-manager-projects/word_counter"
)

// "go_projects/cli-projects/calculator"
// evenoddchecker "go_projects/cli-projects/even_odd_checker"
// "go_projects/cli-projects/fibonacci"

func main() {
	// calculator.Init()
	// evenoddchecker.Check(5)
	// fibonacci.Init()
	scrapeFile()
	reverse()
	copy()
}

func scrapeFile() {
	fileInfo := wordcounter.WordCount("cli-projects/calculator/add.go")
	fmt.Printf("Lines: %d, Words: %d", fileInfo["lines"], fileInfo["words"])
}

func reverse() {
	text := "Hello, World!"
	reversed := textreverser.ReverseText(text)
	fmt.Println("Original:", text)
	fmt.Println("Reversed:", reversed)
}

func copy() {
	// Specify the source and destination file paths
	source := "source.txt"
	destination := "destination.txt"

	// Call the copyFile function
	err := filecopier.CopyFile(source, destination)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File copied successfully!")
	}
}