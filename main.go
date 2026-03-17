package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

func main() {

	clearScreen()

	pathToFile := getInputFilePath()
	fileName := GetFileName(pathToFile)

	fileLines := TryReadFileContent(pathToFile)

	linesBuffer := NewLinesBuffer(fileLines)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	EnterAlternateScreen()
	defer ExitAlternateScreen()

	keyCh := make(chan []byte)

	// Goroutine that blocks on stdin and sends each byte to the channel
	go func() {
		buf := make([]byte, 8)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil || n == 0 {
				close(keyCh)
				return
			}
			tmp := make([]byte, n) // in order to not get overwritten data
			copy(tmp, buf[:n])
			keyCh <- tmp
		}
	}()

	refreshInterval := 300 * time.Millisecond
	ticker := time.NewTicker(refreshInterval)
	defer ticker.Stop()

	initialRenderError := ClearAndRenderScreen(linesBuffer, fileName)
	if initialRenderError != nil {
		LogError(fmt.Sprintf("ERROR: %v", initialRenderError))
		return
	}

	for {
		select {
		case keyBytes, ok := <-keyCh:
			if !ok {
				return
			}
			key := DetermineKey(keyBytes)

			switch key {

			case "KeyCtrlC":
				return
			case "KeyCtrlS":
				err := SaveFile(fileName, linesBuffer.Lines)
				if err != nil {
					LogError(fmt.Sprintf("Error while saving file: %v", err))
				}
				return

			case "KeyBackspace":
				linesBuffer.DeleteCharacterBackward()
			case "KeyDelete":
				linesBuffer.DeleteCharacterForward()

			case "KeyEnter":
				linesBuffer.BreakLine()

			case "KeyUp":
				linesBuffer.MoveCursorUp()
			case "KeyDown":
				linesBuffer.MoveCursorDown()
			case "KeyRight":
				linesBuffer.MoveCursorRight()
			case "KeyLeft":
				linesBuffer.MoveCursorLeft()

			case "KeyEnd":
				linesBuffer.MoveCursorToEndOfLine()
			case "KeyHome":
				linesBuffer.MoveCursorToBeginingOfLine()

			default:
				if len(key) == 1 {
					linesBuffer.InsertPrintableCharacter(key)
				}
			}

			linesBuffer.SetCursorVisibility(true)
			ticker.Reset(refreshInterval) // Reset the ticker so cursor doesn't flicker right after a keypress

			err := ClearAndRenderScreen(linesBuffer, fileName)
			if err != nil {
				LogError(fmt.Sprintf("ERROR: %v", err))
				return
			}

		case <-ticker.C:
			linesBuffer.ToggleCursorVisibilityVisibility()

			err := ClearAndRenderScreen(linesBuffer, fileName)
			if err != nil {
				LogError(fmt.Sprintf("ERROR: %v", err))
				return
			}
		}
	}
}

func getInputFilePath() string {
	var pathToFile string

	if len(os.Args) < 2 {
		fmt.Print("Enter file path: ")
		fmt.Scanln(&pathToFile)
	} else {
		pathToFile = os.Args[1]
	}

	if len(pathToFile) == 0 || pathToFile == "" {
		LogError("No file path provided")
		getInputFilePath()
	}

	return pathToFile
}
