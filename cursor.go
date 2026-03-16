package main

import "fmt"

type Cursor struct {
	LineIndex      int
	CharacterIndex int
	CurrentSymbol  string
	IsVisible      bool
}

func NewCursor() Cursor {
	return Cursor{
		LineIndex:      0,
		CharacterIndex: 0,
		CurrentSymbol:  "|",
		IsVisible:      true,
	}
}

func (cursor *Cursor) getSymbolAccourdingVisibility() string {
	if cursor.IsVisible {
		return fmt.Sprintf("\033[31m%s\033[0m", cursor.CurrentSymbol)
	} else {
		return " "
	}
}

func (cursor *Cursor) SetPosition(lineIndex int, characterIndex int) {
	cursor.LineIndex = lineIndex
	cursor.CharacterIndex = characterIndex
}
