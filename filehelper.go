package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func TryReadFileContent(filePath string) []string {
	result := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		result = append(result, "")
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if(len(result) == 0){
		result = append(result, "")
	}
	
	return result
}

func SaveFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error while creating file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("Error while writing to file: %v", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Error while flushing writer: %v", err)
	}

	return nil
}

func GetFileName(filePath string) string {
	return filepath.Base(filePath)
}
