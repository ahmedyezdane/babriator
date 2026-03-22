package logging

import "fmt"

func LogError(msg string) {
	fmt.Printf("\033[31m[ERROR]\033[0m %s\n", msg)
}

func LogWarn(msg string) {
	fmt.Printf("\033[33m[WARN]\033[0m %s\n", msg)
}
