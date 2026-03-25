package buffer

import "babritor/helpers"

func (buffer *LinesBuffer) InsertPrintableCharacter(content string) {
	if len(content) == 0 || content == "" {
		return
	}

	textLine := ((*buffer).Lines)[buffer.Cursor.LineIndex]

	substring1, substring2 := helpers.SplitStringAtIndex(textLine, buffer.Cursor.CharacterIndex)

	((*buffer).Lines)[buffer.Cursor.LineIndex] = substring1 + content + substring2

	buffer.MoveCursorRight()
	buffer.markLineUpdated(buffer.Cursor.LineIndex)
}

func (buffer *LinesBuffer) DeleteCharacterBackward() {

	if buffer.Cursor.CharacterIndex == 0 {
		if buffer.Cursor.LineIndex == 0 {
			return
		}

		// add current line to the end of previous line
		precedingLines := ((*buffer).Lines)[:(buffer.Cursor.LineIndex - 1)]
		previousLine := ((*buffer).Lines)[buffer.Cursor.LineIndex-1]

		currentLine := ((*buffer).Lines)[buffer.Cursor.LineIndex]

		nextLines := ((*buffer).Lines)[(buffer.Cursor.LineIndex + 1):]

		linesCopy := []string{}

		linesCopy = append(linesCopy, precedingLines...)
		linesCopy = append(linesCopy, previousLine+currentLine)
		linesCopy = append(linesCopy, nextLines...)

		(*buffer).Lines = linesCopy

		buffer.Cursor.SetPosition((buffer.Cursor.LineIndex - 1), len(previousLine))

		buffer.markLineUpdated(buffer.Cursor.LineIndex)
		buffer.markLineRemoved(buffer.Cursor.LineIndex + 1)

	} else {
		textLine := ((*buffer).Lines)[buffer.Cursor.LineIndex]

		substring1, substring2 := helpers.SplitStringAtIndex(textLine, buffer.Cursor.CharacterIndex)

		substring1 = substring1[:len(substring1)-1]

		((*buffer).Lines)[buffer.Cursor.LineIndex] = substring1 + substring2

		buffer.MoveCursorLeft()
		buffer.markLineUpdated(buffer.Cursor.LineIndex)
	}
}

func (buffer *LinesBuffer) DeleteCharacterForward() {

	currentLine := ((*buffer).Lines)[buffer.Cursor.LineIndex]

	isAtTheEndOfCurrentLine := (buffer.Cursor.CharacterIndex == len(currentLine))

	if isAtTheEndOfCurrentLine {

		precedingLines := ((*buffer).Lines)[:buffer.Cursor.LineIndex]

		nextLine := ((*buffer).Lines)[buffer.Cursor.LineIndex+1]

		followingLines := ((*buffer).Lines)[(buffer.Cursor.LineIndex + 2):]

		linesCopy := []string{}

		linesCopy = append(linesCopy, precedingLines...)
		linesCopy = append(linesCopy, currentLine+nextLine)
		linesCopy = append(linesCopy, followingLines...)

		(*buffer).Lines = linesCopy
		buffer.markLineUpdated(buffer.Cursor.LineIndex)
		buffer.markLineRemoved(buffer.Cursor.LineIndex + 1)
	} else {
		substring1, substring2 := helpers.SplitStringAtIndex(currentLine, buffer.Cursor.CharacterIndex)

		substring2 = substring2[1:]

		((*buffer).Lines)[buffer.Cursor.LineIndex] = substring1 + substring2
		buffer.markLineUpdated(buffer.Cursor.LineIndex)
	}
}

func (buffer *LinesBuffer) BreakLine() {
	previousLines := ((*buffer).Lines)[:buffer.Cursor.LineIndex]
	currentLine := ((*buffer).Lines)[buffer.Cursor.LineIndex]
	nextLines := ((*buffer).Lines)[(buffer.Cursor.LineIndex + 1):]

	linesCopy := []string{}

	linesCopy = append(linesCopy, previousLines...)

	substring1, substring2 := helpers.SplitStringAtIndex(currentLine, buffer.Cursor.CharacterIndex)
	linesCopy = append(linesCopy, substring1)
	linesCopy = append(linesCopy, substring2)

	linesCopy = append(linesCopy, nextLines...)

	(*buffer).Lines = linesCopy

	buffer.MoveCursorRight()
	buffer.markLineUpdated(buffer.Cursor.LineIndex - 1)
	buffer.markLineCreated(buffer.Cursor.LineIndex)
}