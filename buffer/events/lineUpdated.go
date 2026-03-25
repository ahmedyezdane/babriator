package events

type LineUpdatedEvent struct {
	LineIndex int
}

func (LineUpdatedEvent) isLineBufferEvent() {}