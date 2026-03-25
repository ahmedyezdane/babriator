package screen

import (
	"babritor/buffer"
	"babritor/buffer/events"
	"fmt"
	"os"
	"strings"
)

// \033[32m - Sets text color to green
// \033[33m - Sets text color to yellow
// \033[0m  - Resets formatting to default
// %-*d     - Left-aligned integer with dynamic width (the * means width is passed as an argument)
const lineNumberTemplate string = "\033[32m%-*d\033[0m \033[33m|\033[0m %v\n"

const (
	headerLinesCount                int8 = 2
	minLinePadding                  int8 = 3
	graphicalCharactersPerLineCount int8 = 4
)

type ScreenState struct {
	headerLinesCount       int8
	paddingCharactersCount int8

	FileName string
}

func NewScreenState(fileName string) ScreenState {
	return ScreenState{
		headerLinesCount:       headerLinesCount,
		paddingCharactersCount: minLinePadding,
		FileName:               fileName,
	}
}

func InitialScreenRender(linesBuffer buffer.LinesBuffer, screenState ScreenState) {

	ClearScreen()

	fmt.Printf("[%v]\n\n", strings.ToUpper(screenState.FileName))

	for i, t := range linesBuffer.Lines {
		fmt.Print(createLineForPrint(screenState.paddingCharactersCount, i+1, t))
	}

	RelocateCursor(linesBuffer, screenState)
}

func ClearScreen() {
	fmt.Print("\033[H\033[3J\033[2J")
}

func createLineForPrint(padding int8, lineNumber int, content string) string {
	return fmt.Sprintf(lineNumberTemplate, padding, lineNumber, content)
}

func RelocateCursor(linesBuffer buffer.LinesBuffer, screenState ScreenState) {
	lineIndex, characterIndex := calculateTerminalCursorPostion(screenState, linesBuffer)

	moveCursorTo(lineIndex, characterIndex)
}

func EnterAlternateScreen() {
	fmt.Print(ansiEnterAlternateScreen)
}

func ExitAlternateScreen() {
	fmt.Print(ansiExitAlternateScreen)
}

func (screenState *ScreenState) ApplyLineBufferEvents(linesBuffer *buffer.LinesBuffer) {
	fmt.Fprint(os.Stdout, ansiHideCursor)

	for _, event := range linesBuffer.Events {

		switch ev := event.(type) {

		case events.LineCreatedEvent:
			screenState.rerenderLinesTillTheEnd(ev.LineIndex, *linesBuffer)

		case events.LineUpdatedEvent:
			screenState.rerenderLine(ev.LineIndex, linesBuffer.Lines[ev.LineIndex])

		case events.LineRemovedEvent:
			screenState.rerenderLinesTillTheEnd(ev.LineIndex, *linesBuffer)
			screenState.clearLastLine(len(linesBuffer.Lines))
		}
	}

	RelocateCursor(*linesBuffer, *screenState)

	linesBuffer.ClearChangeEvents()

	fmt.Fprint(os.Stdout, ansiShowCursor)
}

func moveCursorTo(lineIndex, characterIndex int) {
	fmt.Printf(ansiMoveCursor, lineIndex, characterIndex)
}

func calculateTerminalCursorPostion(screenState ScreenState, linesBuffer buffer.LinesBuffer) (lineNumber, characterIndex int) {

	lineNumber = int(screenState.headerLinesCount) +
		linesBuffer.Cursor.LineIndex +
		1 //terminal line number start at 1 rather than 0

	characterIndex = int(screenState.paddingCharactersCount) +
		int(graphicalCharactersPerLineCount) +
		linesBuffer.Cursor.CharacterIndex

	return
}

func calculateTerminalCursorLineNumber(screenState ScreenState, cursorLineIndex int) int {
	return int(screenState.headerLinesCount) +
		cursorLineIndex +
		1
}

func (screenState *ScreenState) rerenderLinesTillTheEnd(startIndex int, linesBuffer buffer.LinesBuffer) {
	for i := startIndex; i < len(linesBuffer.Lines); i++ {
		screenState.rerenderLine(i, linesBuffer.Lines[i])
	}
}

func (screenState *ScreenState) rerenderLine(lineIndex int, text string) {
	tln := calculateTerminalCursorLineNumber(*screenState, lineIndex)
	moveCursorTo(tln, 0)
	fmt.Fprint(os.Stdout, ansiClearLine)
	lineContent := createLineForPrint(screenState.paddingCharactersCount, lineIndex+1, text)
	fmt.Fprint(os.Stdout, lineContent)
}

func (screenState *ScreenState) clearLastLine(lastLineIndex int) {
	tln := calculateTerminalCursorLineNumber(*screenState, lastLineIndex)
	moveCursorTo(tln, 0)
	fmt.Fprint(os.Stdout, ansiDeleteLine)
}
