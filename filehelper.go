package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileContent(filename string) ([]string, error) {

	result := []string{}

	file, err := os.Open(filename)
	if err != nil {
		return result, fmt.Errorf("Error while opening file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	err = scanner.Err()
	if err != nil {
		return []string{}, fmt.Errorf("Error while scanning file: %v", err)
	}

	return result, nil
}
