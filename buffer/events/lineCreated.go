package events

type LineCreatedEvent struct {
	LineIndex int
}

func (LineCreatedEvent) isLineBufferEvent() {}