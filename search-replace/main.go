package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchAndReplace(filename, search, replace string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), search, replace, -1)
		lines = append(lines, line)
	}

	return ReplaceText(filename, lines)
}

func ReplaceText(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	argLen := len(os.Args[1:])
	if argLen < 3 {
		fmt.Println("Filename, search, replace are needed")
		return
	}

	filename := os.Args[1]
	search   := os.Args[2]
	replace  := os.Args[3]

	SearchAndReplace(filename, search, replace)
}