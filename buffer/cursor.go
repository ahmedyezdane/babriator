package buffer

type Cursor struct {
	LineIndex      int
	CharacterIndex int
}

func NewCursor() Cursor {
	return Cursor{
		LineIndex:      0,
		CharacterIndex: 0,
	}
}

func (cursor *Cursor) SetPosition(lineIndex int, characterIndex int) {
	cursor.LineIndex = lineIndex
	cursor.CharacterIndex = characterIndex
}
