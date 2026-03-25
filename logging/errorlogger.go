package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorGray   = "\033[90m"
)

var (
	logFile *os.File
	logger  *log.Logger
	enabled = true
)

func Initialize() error {
	if !enabled {
		return nil
	}

	logPath := filepath.Join(os.TempDir(), "babriator.log")

	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("\n\n\n%s=== Editor started ===%s\n\n\n",ColorGreen,ColorReset)

	return nil
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}

func LogError(msg string) {
	if !enabled || logger == nil {
		return
	}
	logger.Printf("%s[ERROR]%s %s\n", ColorRed, ColorReset, msg)
}

func LogWarn(msg string) {
	if !enabled || logger == nil {
		return
	}
	logger.Printf("%s[WARN]%s %s\n", ColorYellow, ColorReset, msg)
}

func LogInfo(msg string) {
	if !enabled || logger == nil {
		return
	}
	logger.Printf("%s[INFO]%s %s\n", ColorBlue, ColorReset, msg)
}

func LogDebug(format string, args ...interface{}) {
	if !enabled || logger == nil {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.Printf("%s[DEBUG]%s %s\n", ColorGray, ColorReset, msg)
}
