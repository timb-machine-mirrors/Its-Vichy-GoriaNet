package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Debug(Content string) {
	if DebugEnabled {
		fmt.Println(Content)
	}
}

func HandleError(err error) bool {
	if err != nil {
		Debug(err.Error())
		return true
	}

	return false
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func InListInt(in int, list []int) bool {
	for _, v := range list {
		if in == v {
			return true
		}
	}
	return false
}

// check if folder exist
func FolderExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
