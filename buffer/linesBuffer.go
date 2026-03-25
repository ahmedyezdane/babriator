package buffer

import "babritor/buffer/events"

type LinesBuffer struct {
	Lines  []string
	Cursor Cursor
	Events []LineBufferEvent
}

func NewLinesBuffer(lines []string) LinesBuffer {
	return LinesBuffer{
		Lines:  lines,
		Cursor: NewCursor(),
		Events: make([]LineBufferEvent, 0),
	}
}

func (lb *LinesBuffer) markLineCreated(lineIndex int) {
	event := events.LineCreatedEvent{
		LineIndex: lineIndex,
	}

	lb.Events = append(lb.Events, event)
}

func (lb *LinesBuffer) markLineUpdated(lineIndex int) {
	event := events.LineUpdatedEvent{
		LineIndex: lineIndex,
	}

	lb.Events = append(lb.Events, event)
}

func (lb *LinesBuffer) markLineRemoved(lineIndex int) {
	event := events.LineRemovedEvent{
		LineIndex: lineIndex,
	}

	lb.Events = append(lb.Events, event)
}

func (lb *LinesBuffer) ClearChangeEvents() {
	lb.Events = lb.Events[:0]
}
