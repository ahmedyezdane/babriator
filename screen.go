package main

import (
	"fmt"
	"strings"
)

func ClearAndRenderScreen(linesBuffer LinesBuffer, fileName string) {
	clearScreen()

	marginLinesCount, marginCharactersCount := renderAndCalculateMargin(linesBuffer.Lines, linesBuffer.Cursor, fileName)

	moveCursorTo(
		linesBuffer.Cursor.LineIndex+marginLinesCount,
		linesBuffer.Cursor.CharacterIndex+marginCharactersCount)

}

func clearScreen() {
	fmt.Print("\033[H\033[3J\033[2J")
}

func renderAndCalculateMargin(textLines []string, cursor Cursor, fileName string) (marginLinesCount, marginCharactersCount int) {

	fmt.Printf("\n[%v]\n\n", strings.ToUpper(fileName))
	marginLinesCount = 4

	marginAccourdingTotalLinesCount := len(fmt.Sprintf("%d", len(textLines)))

	for i, line := range textLines {
		fmt.Printf("\033[32m%-*d\033[0m \033[33m|\033[0m %v\n", marginAccourdingTotalLinesCount, i+1, line)
	}

	marginCharactersCount = marginAccourdingTotalLinesCount + 4

	return
}

func EnterAlternateScreen() {
	fmt.Print("\033[?1049h")
}

func ExitAlternateScreen() {
	fmt.Print("\033[?1049l")
}

const (
	// Move cursor to position (row, col) - both 1-based
	ansiMoveCursor = "\033[%d;%dH"

	// Cursor style (DECSCUSR)
	ansiCursorBlinkingBlock     = "\033[1 q"
	ansiCursorSteadyBlock       = "\033[2 q"
	ansiCursorBlinkingUnderline = "\033[3 q"
	ansiCursorSteadyUnderline   = "\033[4 q"
	ansiCursorBlinkingBar       = "\033[5 q" // classic insert-mode cursor
	ansiCursorSteadyBar         = "\033[6 q"
)

func moveCursorTo(row, col int) {
	fmt.Printf(ansiMoveCursor, row, col)
}
