package main

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"babritor/buffer"
	"babritor/helpers"
	"babritor/logging"
	"babritor/screen"
)

const VERSION string = "1.1.0"

func main() {
	
	if err := logging.Initialize(); err != nil {
		panic(err)
	}
	defer logging.Close()

	screen.ClearScreen()

	pathToFile := getInputFilePath()
	fileName := helpers.GetFileName(pathToFile)

	fileLines := helpers.TryReadFileContent(pathToFile)

	linesBuffer := buffer.NewLinesBuffer(fileLines)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	screen.EnterAlternateScreen()
	defer screen.ExitAlternateScreen()

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

	screenState := screen.NewScreenState(fileName)
	screen.InitialScreenRender(linesBuffer, screenState)

	for keyBytes := range keyCh {
		key := helpers.DetermineKey(keyBytes)

		switch key {

		case "KeyCtrlC":
			return
		case "KeyCtrlS":
			err := helpers.SaveFile(fileName, linesBuffer.Lines)
			if err != nil {
				logging.LogError(fmt.Sprintf("Error while saving file: %v", err))
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

		screenState.ApplyLineBufferEvents(&linesBuffer)
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
		logging.LogError("No file path provided")
		getInputFilePath()
	}

	return pathToFile
}
