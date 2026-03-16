package main

import (
	"fmt"
	"strings"
)

func ClearAndRenderScreen(linesBuffer LinesBuffer, fileName string) error {
	clearScreen()
	err := render(linesBuffer.Lines, linesBuffer.Cursor, fileName)
	return err
}

func clearScreen() {
    fmt.Print("\033[H\033[3J\033[2J")
}


func render(textLines []string, cursor Cursor, fileName string) error {

	printableLines, err := appendCursorToTextLines(textLines, cursor)
	if err != nil {
		return err
	}

	fmt.Printf("\n[%v]\n\n",strings.ToUpper(fileName))

	for i, line := range printableLines {
		fmt.Printf("\033[32m%-*d\033[0m \033[33m|\033[0m %v\n", len(fmt.Sprintf("%d", len(printableLines))), i+1, line)
	}

	return nil
}

func appendCursorToTextLines(textLines []string, cursor Cursor) ([]string, error) {
	textLinesCount := len(textLines)

	output := make([]string, textLinesCount)

	// TODO<Ahmad> test this case
	if textLinesCount == 0 {
		output = append(output, cursor.getSymbolAccourdingVisibility())
		return output, nil
	}

	copy(output, textLines)

	textLine := output[cursor.LineIndex]

	substring1 := textLine[:cursor.CharacterIndex]
	substring2 := textLine[cursor.CharacterIndex:]

	output[cursor.LineIndex] = substring1 + cursor.getSymbolAccourdingVisibility() + substring2

	return output, nil
}

func enterAlternateScreen() {
    fmt.Print("\033[?1049h")
}

func exitAlternateScreen() {
    fmt.Print("\033[?1049l")
}
