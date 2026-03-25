package buffer

func (buffer *LinesBuffer) MoveCursorUp() {
	if buffer.Cursor.LineIndex > 0 {
		buffer.Cursor.LineIndex--
		// clamp CharacterIndex to new line length
		lineLen := len(buffer.Lines[buffer.Cursor.LineIndex])
		if buffer.Cursor.CharacterIndex > lineLen {
			buffer.Cursor.CharacterIndex = lineLen
		}
	}
}

func (buffer *LinesBuffer) MoveCursorDown() {
	if buffer.Cursor.LineIndex < len(buffer.Lines)-1 {
		buffer.Cursor.LineIndex++
		lineLen := len(buffer.Lines[buffer.Cursor.LineIndex])
		if buffer.Cursor.CharacterIndex > lineLen {
			buffer.Cursor.CharacterIndex = lineLen
		}
	}
}

func (buffer *LinesBuffer) MoveCursorLeft() {
	if buffer.Cursor.CharacterIndex > 0 {
		buffer.Cursor.CharacterIndex--
	} else if buffer.Cursor.LineIndex > 0 {
		buffer.Cursor.LineIndex--
		buffer.Cursor.CharacterIndex = len(buffer.Lines[buffer.Cursor.LineIndex])
	}
}

func (buffer *LinesBuffer) MoveCursorRight() {
	lineLen := len(buffer.Lines[buffer.Cursor.LineIndex])
	if buffer.Cursor.CharacterIndex < lineLen {
		buffer.Cursor.CharacterIndex++
	} else if buffer.Cursor.LineIndex < len(buffer.Lines)-1 {
		buffer.Cursor.LineIndex++
		buffer.Cursor.CharacterIndex = 0
	}
}

func (buffer *LinesBuffer) MoveCursorToEndOfLine() {
	lineLen := len(buffer.Lines[buffer.Cursor.LineIndex])
	buffer.Cursor.CharacterIndex = lineLen
}

func (buffer *LinesBuffer) MoveCursorToBeginingOfLine() {
	buffer.Cursor.CharacterIndex = 0
}
