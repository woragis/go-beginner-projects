package wordcounter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func WordCount(fileName string) (map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	lineCount := 0
	wordCount := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++

		words:= strings.Fields(scanner.Text())
		wordCount+=len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}


	return map[string]int{
		"lines": lineCount,
		"words": wordCount,
	}
}