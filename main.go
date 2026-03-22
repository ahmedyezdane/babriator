package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

const VERSION string = "1.1.0"

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

	go func() {
		buf := make([]byte, 8)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil || n == 0 {
				close(keyCh)
				return
			}
			tmp := make([]byte, n)
			copy(tmp, buf[:n])
			keyCh <- tmp
		}
	}()

	ClearAndRenderScreen(linesBuffer, fileName)

	for keyBytes := range keyCh {
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

		ClearAndRenderScreen(linesBuffer, fileName)
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
