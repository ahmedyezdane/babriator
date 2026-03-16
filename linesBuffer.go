package main

type LinesBuffer struct {
	Lines  []string
	Cursor Cursor
}

func NewLinesBuffer(lines []string) LinesBuffer {
	return LinesBuffer{
		Lines:  lines,
		Cursor: NewCursor(),
	}
}

func (b *LinesBuffer) MoveCursorUp() {
	if b.Cursor.LineIndex > 0 {
		b.Cursor.LineIndex--
		// clamp CharacterIndex to new line length
		lineLen := len(b.Lines[b.Cursor.LineIndex])
		if b.Cursor.CharacterIndex > lineLen {
			b.Cursor.CharacterIndex = lineLen
		}
	}
}

func (b *LinesBuffer) MoveCursorDown() {
	if b.Cursor.LineIndex < len(b.Lines)-1 {
		b.Cursor.LineIndex++
		lineLen := len(b.Lines[b.Cursor.LineIndex])
		if b.Cursor.CharacterIndex > lineLen {
			b.Cursor.CharacterIndex = lineLen
		}
	}
}

func (b *LinesBuffer) MoveCursorLeft() {
	if b.Cursor.CharacterIndex > 0 {
		b.Cursor.CharacterIndex--
	} else if b.Cursor.LineIndex > 0 {
		b.Cursor.LineIndex--
		b.Cursor.CharacterIndex = len(b.Lines[b.Cursor.LineIndex])
	}
}

func (b *LinesBuffer) MoveCursorRight() {
	lineLen := len(b.Lines[b.Cursor.LineIndex])
	if b.Cursor.CharacterIndex < lineLen {
		b.Cursor.CharacterIndex++
	} else if b.Cursor.LineIndex < len(b.Lines)-1 {
		b.Cursor.LineIndex++
		b.Cursor.CharacterIndex = 0
	}
}

func (b *LinesBuffer) MoveCursorToEndOfLine() {
	lineLen := len(b.Lines[b.Cursor.LineIndex])
	b.Cursor.CharacterIndex = lineLen
}
func (b *LinesBuffer) MoveCursorToBeginingOfLine() {
	b.Cursor.CharacterIndex = 0
}

func (b *LinesBuffer) SetCursorVisibility(isVisibile bool) {
	b.Cursor.IsVisible = isVisibile
}
func (b *LinesBuffer) ToggleCursorVisibilityVisibility() {
	b.Cursor.IsVisible = !b.Cursor.IsVisible
}

func (linesBuffer *LinesBuffer) InsertPrintableCharacter(content string) {
	if len(content) == 0 || content == "" {
		return
	}

	textLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex]

	substring1, substring2 := splitLine(textLine, linesBuffer.Cursor.CharacterIndex)

	((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex] = substring1 + content + substring2

	linesBuffer.MoveCursorRight()
}

func (linesBuffer *LinesBuffer) DeleteCharacterBackward() {

	if linesBuffer.Cursor.CharacterIndex == 0 {
		// add current line to the end of previous line

		precedingLines := ((*linesBuffer).Lines)[:(linesBuffer.Cursor.LineIndex - 1)]
		previousLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex-1]

		currentLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex]

		nextLines := ((*linesBuffer).Lines)[(linesBuffer.Cursor.LineIndex + 1):]

		linesCopy := []string{}

		linesCopy = append(linesCopy, precedingLines...)
		linesCopy = append(linesCopy, previousLine+currentLine)
		linesCopy = append(linesCopy, nextLines...)

		(*linesBuffer).Lines = linesCopy

		linesBuffer.Cursor.SetPosition((linesBuffer.Cursor.LineIndex - 1), len(previousLine))

	} else {
		textLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex]

		substring1, substring2 := splitLine(textLine, linesBuffer.Cursor.CharacterIndex)

		substring1 = substring1[:len(substring1)-1]

		((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex] = substring1 + substring2

		linesBuffer.MoveCursorLeft()
	}
}

func (linesBuffer *LinesBuffer) DeleteCharacterForward() {

	currentLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex]

	isAtTheEndOfCurrentLine := (linesBuffer.Cursor.CharacterIndex == len(currentLine))

	if isAtTheEndOfCurrentLine {

		precedingLines := ((*linesBuffer).Lines)[:linesBuffer.Cursor.LineIndex]

		nextLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex+1]

		followingLines := ((*linesBuffer).Lines)[(linesBuffer.Cursor.LineIndex + 2):]

		linesCopy := []string{}

		linesCopy = append(linesCopy, precedingLines...)
		linesCopy = append(linesCopy, currentLine+nextLine)
		linesCopy = append(linesCopy, followingLines...)

		(*linesBuffer).Lines = linesCopy

		//linesBuffer.Cursor.SetPosition((linesBuffer.Cursor.LineIndex - 1), len(previousLine))
	} else {
		substring1, substring2 := splitLine(currentLine, linesBuffer.Cursor.CharacterIndex)

		substring2 = substring2[1:]

		((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex] = substring1 + substring2
	}
}

func (linesBuffer *LinesBuffer) BreakLine() {
	previousLines := ((*linesBuffer).Lines)[:linesBuffer.Cursor.LineIndex]
	currentLine := ((*linesBuffer).Lines)[linesBuffer.Cursor.LineIndex]
	nextLines := ((*linesBuffer).Lines)[(linesBuffer.Cursor.LineIndex + 1):]

	linesCopy := []string{}

	linesCopy = append(linesCopy, previousLines...)

	substring1, substring2 := splitLine(currentLine, linesBuffer.Cursor.CharacterIndex)
	linesCopy = append(linesCopy, substring1)
	linesCopy = append(linesCopy, substring2)

	linesCopy = append(linesCopy, nextLines...)

	(*linesBuffer).Lines = linesCopy

	linesBuffer.MoveCursorRight()
}

func splitLine(textLine string, characterIndex int) (string, string) {
	substring1 := textLine[:characterIndex]
	substring2 := textLine[characterIndex:]

	return substring1, substring2
}
