package events

type LineRemovedEvent struct {
	LineIndex int
}

func (LineRemovedEvent) isLineBufferEvent() {}
